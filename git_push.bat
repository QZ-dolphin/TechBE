@echo off

REM 获取当前日期和时间，并格式化为 YYYY-MM-DD_HH-MM-SS
for /f "tokens=1-5 delims=/: " %%a in ('echo %date% %time%') do (
    set datetime=%%c-%%a-%%b_%%d-%%e
)

REM 定义变量a为"revise"加上当前日期时间
set a=revise_%datetime%

REM 添加所有更改到git
git add .
if %ERRORLEVEL% neq 0 (
    echo Failed to add changes. Exiting.
    exit /b %ERRORLEVEL%
)


REM 使用变量a作为提交信息
git commit -m "%a%"
if %ERRORLEVEL% neq 0 (
    echo Failed to commit changes. Exiting.
    exit /b %ERRORLEVEL%
)

REM 推送更改到master分支
git push origin master
if %ERRORLEVEL% neq 0 (
    echo Failed to push changes. Exiting.
    exit /b %ERRORLEVEL%
)

echo All git commands executed successfully.
