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
![](https://raw.githubusercontent.com/channyein1337/gup/main/gup.png)

### For Single Url
For a single parameter
```
▶ echo "http://localhost/index.php?id=1&redirect=true" | gup -s
```
![](https://github.com/channyein1337/gup/blob/main/gup-single-param.png?raw=true)
For multiple parameters
```
▶ echo "http://localhost/index.php?id=1&redirect=true" | gup -m
```
![](https://github.com/channyein1337/gup/blob/main/gup_multiple.png?raw=true)

### For Mutliple Urls
For a single parameter
```
cat gau.txt | gup -s | sort -u
```
![](https://github.com/channyein1337/gup/blob/main/gup-single-file.png?raw=true)
For multiple parameters
```
cat gau.txt | gup -m | sort -u 
```
![](https://github.com/channyein1337/gup/blob/main/gup-multi-file.png?raw=true)
