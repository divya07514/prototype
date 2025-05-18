# 🧠 Append-Only Key-Value Store (Go)

This is a simple **key-value store** implemented in Go using:

- An **append-only file** to store all writes (values + deletes)
- An **in-memory index** for fast lookups (maps keys to file offsets)
- A **tombstone mechanism** for deletes (no in-place updates)

---

## 📦 Features

- ✅ **Persistent log**: All writes are appended to disk
- ✅ **In-memory index**: Fast lookup using a Go `map[string]int64`
- ✅ **Delete support** using tombstones (soft deletes)

---

## 🚀 Basic Usage

```go
kv := NewKVStore("data.db")

kv.Put("foo", []byte("bar"))

value, err := kv.Get("foo")
// value = []byte("bar")

kv.Delete("foo")

value, err := kv.Get("foo")
// err = "key was deleted"
```

---

## 🧱 On-Disk Entry Format

Each record in the file is:

| Field       | Type     | Size        | Description                    |
|-------------|----------|-------------|--------------------------------|
| `keyLen`    | `int32`  | 4 bytes     | Length of the key              |
| `valLen`    | `int32`  | 4 bytes     | Length of the value            |
| `tombstone` | `bool`   | 1 byte      | `false` = live, `true` = deleted |
| `key`       | `[]byte` | keyLen bytes| Key bytes                      |
| `value`     | `[]byte` | valLen bytes| Value bytes (if not deleted)   |

---

## 🧠 How It Works

- `Put(key, value)`:
  - Appends a new entry with tombstone = false
  - Updates the in-memory index to the new offset

- `Delete(key)`:
  - Appends an entry with tombstone = true and no value
  - Removes the key from the in-memory index

- `Get(key)`:
  - Looks up the offset from the in-memory index
  - Reads the record, checks tombstone
  - Returns the value if not deleted

---

## 🛠️ What You Can Build On Top

Here are some extensions and ideas:

### 🔄 Compaction
- Periodically rewrite a clean file with only the latest non-deleted entries
- Make it a background task to avoid blocking

### ⏱️ TTL Support (Time to Live)
- Add timestamp or expiry time in the record
- Garbage collect expired keys during compaction or read

### 🧵 Concurrency
- Add `sync.RWMutex` for thread-safe access
- Goroutine-safe read and write APIs

### 🧮 Value Versioning
- Store version numbers or timestamps in the entry
- Allow multi-version concurrency control or auditing

### 📦 Serialization Formats
- Add support for structured values (e.g., JSON, protobuf)

### 📦 Rebuilt Index
- Add support for rebuilding the in-memory index from the file in case of a crash
