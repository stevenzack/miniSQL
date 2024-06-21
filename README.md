# miniSQL
A minimal SQL database written in Go


Our B-tree will be persisted to the disk eventually, so we need to design the wire format for the B-tree nodes first. Without the format, we won’t know the size of a node and when to split a node.
A node consists of:
1. A fixed-sized header containing the type of the node (leaf node or internal node) and the number of keys.
2. A list of pointers to the child nodes. (Used by internal nodes).
3. A list of offsets pointing to each key-value pair.
4. Packed KV pairs.

| type | nkeys | pointers | offsets | key-values |
| - | - | - | - | - |
| 2B | 2B | nkeys*8B | nkeys*2B | .. |

This is the format of the KV pair. Lengths followed by data.

| klen | vlen | key | val |
| - | - | - | - |
| 2B | 2B | <1000 | <3000 |

To keep things simple, both leaf nodes and internal nodes use the same format.