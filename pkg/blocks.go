// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright (C) 2022  Marten Mooij
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pst

import (
	"encoding/binary"

	"github.com/rotisserie/eris"
)

// GetBlockSize returns the size of a block.
// References "Blocks".
func (file *File) GetBlockSize() (int, error) {
	switch file.FormatType {
	case FormatTypeUnicode:
		return 8192, nil
	case FormatTypeUnicode4k:
		return 65536, nil
	case FormatTypeANSI:
		return 8192, nil
	default:
		return 0, ErrFormatTypeUnsupported
	}
}

// GetBlockTrailerSize returns the size of a block trailer.
// References "Blocks".
func (file *File) GetBlockTrailerSize() (int, error) {
	switch file.FormatType {
	case FormatTypeUnicode:
		return 16, nil
	case FormatTypeUnicode4k:
		return 16, nil
	case FormatTypeANSI:
		return 12, nil
	default:
		return 0, ErrFormatTypeUnsupported
	}
}

// BlockType represents a XBlock or XXBlock.
type BlockType uint8

// Constants defining the block types.
const (
	BlockTypeXBlock  BlockType = 1
	BlockTypeXXBlock BlockType = 2
)

// GetBlocks returns all blocks (XBlock/XXBlock) from a Heap-on-Node along with the total blocks size.
// Internal identifiers have blocks.
//
// References:
// - https://github.com/mooijtech/go-pst/tree/master/docs#xblock
// - https://github.com/mooijtech/go-pst/tree/master/docs#xxblock
func (file *File) GetBlocks(btreeNodeHeapOnNodeOffset int64) ([]BTreeNode, error) {
	data := make([]byte, 4)

	if _, err := file.Reader.ReadAt(data, btreeNodeHeapOnNodeOffset); err != nil {
		return nil, eris.Wrap(err, "failed to read block data")
	}

	blockSignature := data[0]                                // Must indicate 1.
	blockLevel := data[1]                                    // 1 indicates XBlock, 2 indicates XXBlock.
	entryCount := int(binary.LittleEndian.Uint16(data[2:4])) // The number of block b-tree identifiers in this XBlock or XXBlock.

	if blockSignature != 1 {
		return nil, ErrBlockSignatureInvalid
	}

	identifierSize := int(GetIdentifierSize(file.FormatType))

	var blocks []BTreeNode

	switch BlockType(blockLevel) {
	case BlockTypeXBlock:
		// XBlock
		blockIdentifiers := make([]byte, entryCount*identifierSize)

		if _, err := file.Reader.ReadAt(blockIdentifiers, btreeNodeHeapOnNodeOffset+8); err != nil {
			return nil, eris.Wrap(err, "failed to read block identifiers")
		}

		for i := 0; i < entryCount; i++ {
			blockIdentifier := GetIdentifierFromBytes(blockIdentifiers[i*identifierSize:(i*identifierSize)+identifierSize], file.FormatType)
			blockBTreeNode, err := file.GetBlockBTreeNode(blockIdentifier) // TODO - Async then wait for the block b-tree node lookups.

			if err != nil {
				return nil, eris.Wrap(err, "failed to find block b-tree node")
			}

			blocks = append(blocks, blockBTreeNode)
		}
	case BlockTypeXXBlock:
		// XXBlock
		blockIdentifiers := make([]byte, entryCount*int(GetIdentifierSize(file.FormatType)))

		if _, err := file.Reader.ReadAt(blockIdentifiers, btreeNodeHeapOnNodeOffset+8); err != nil {
			return nil, eris.Wrap(err, "failed to read block identifiers")
		}

		for i := 0; i < entryCount; i++ {
			blockIdentifier := GetIdentifierFromBytes(blockIdentifiers[i*identifierSize:(i*identifierSize)+identifierSize], file.FormatType)
			blockBTreeNode, err := file.GetBlockBTreeNode(blockIdentifier) // TODO - Async then wait for the block b-tree node lookups.

			if err != nil {
				return nil, eris.Wrap(err, "failed to find block b-tree node")
			}

			// Recursive.
			blockBTreeNodeBlocks, err := file.GetBlocks(blockBTreeNode.FileOffset)

			if err != nil {
				return nil, eris.Wrap(err, "failed to get blocks")
			}

			blocks = append(blocks, blockBTreeNodeBlocks...)
		}
	default:
		return nil, ErrBlockTypeInvalid
	}

	return blocks, nil
}

// GetBlocksTotalSize returns the size of the external data referenced by the XBlock or XXBlock.
func (file *File) GetBlocksTotalSize(nodeEntryHeapOnNodeOffset int64) (uint32, error) {
	blocksTotalSize := make([]byte, 4)

	if _, err := file.Reader.ReadAt(blocksTotalSize, nodeEntryHeapOnNodeOffset+4); err != nil {
		return 0, eris.Wrap(err, "failed to read total blocks size")
	}

	return binary.LittleEndian.Uint32(blocksTotalSize), nil
}
