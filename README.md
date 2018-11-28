# MyArticles
This is a Go project aimed to provide article management service with Restful API

## install
1. Install glide  
Please get glide package below:  
[glide download link](https://github.com/Masterminds/glide/releases/tag/v0.12.3)  
Notice: Use glide v0.12.3 to avoid vendor installtion issue in windows.
2. Put source code root  
(gopath)/src/github.com/leonkaihao/myarticles  
3. cd into source code root and run install  
$glide install
4. issue with go-sqlite3 on windows  
Users may counter a issue to build app with gcc on windows. If you do, remember to install msys first.  
http://www.msys2.org/  
Then run command below in msys terminal to install gcc:  
$ pacman -S mingw-w64-x86_64-gcc  
Finally add env var to PATH:  
(msys64 folder)\mingw64\bin
4. Build source code  
$go build
## run  
$myarticles  
If user see this info:  
```
2018/11/28 17:34:40 Open database  ./service.db ...  
2018/11/28 17:34:40 Start server :3000 ...  
```
Means server started successfully.

## Test tool
I suggest to use:  
https://www.getpostman.com/
to test restful http request. It is easy to post restful request and has a history list to trace results.

## Unit test
The unit tests are basically focusing on database interfaces.
```
$cd services/database
$go test
```
Recommend vscode to do unit test, its golang plugin associate different go tools to the editor.  
The run test or debug test on file or package has really good experience.

## Development doc
Development doc can be found in (source root)/development_guide.docx
