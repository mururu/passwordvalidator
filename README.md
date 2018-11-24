# passwordvalidator

## Usage

```
passwordvalidator.CommonPassword("dou394l!2i") => true
passwordvalidator.CommonPassword("admin") => false
```

```
passwordvalidator.Similarity("dou394l!2i", "xxxxx@example.com") => true
passwordvalidator.Similarity("xxxxx", "xxxxx@eample.com") => false
```
