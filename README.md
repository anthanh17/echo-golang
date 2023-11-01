# echo-golang

docker exec -it echo-postgres-container bash
psql -d golang -U annt
\dt

-------------
Test postmain
POST: http://localhost:3000/user/sign-up
JSON body:
{
    "Email": "annguyenthe@gmail.com",
    "Password": "abc123",
    "FullName": "Nguyen The An"
}