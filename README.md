# Description
GUP is a tool to create wrodlists from parameters and use those parameters for brute forcing.

# Install
```
▶ go get github.com/channyein1337/gup
```

# Usage
```
▶ gup -h
Usage of gup:
  -m    For mutliple parameters
  -s    For a single parameter
```

### For Single Url
For a single parameter
```
▶ echo "http://localhost/index.php?id=1&redirect=true" | gup -s
```
For multiple parameters
```
▶ echo "http://localhost/index.php?id=1&redirect=true" | gup -m
```

### For Mutliple Urls
For a single parameter
```
cat gau.txt | gup -s | sort -u
```
For multiple parameters
```
cat gau.txt | gup -m | sort -u 
```
