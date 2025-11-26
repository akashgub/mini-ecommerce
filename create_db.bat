@echo off
REM Set PostgreSQL password
set PGPASSWORD=4152
REM Create database
"C:\Program Files\PostgreSQL\15\bin\createdb.exe" -U postgres ecommerce
echo Database created successfully!
pause
