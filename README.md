SMNotes                       
Created by: JustSMN

Welcome to my golang notes web application!  
This application uses http handles to create, view, and delete notes, and it saves all notes into a postgresql database.

How to Use:
1. Run the PostgreSQL with Docker:\
   `PS C:\Programming\smnotes> docker pull postgres` \
   `PS C:\Programming\smnotes> docker run --name web-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres` \
   (For windows)
    
2. Download the libraries:\
   `PS C:\Programming\smnotes> go mod download`

3. Run main.go in terminal:\
   `PS C:\Programming\ai-tg-bot> go run main.go`

4. Go to `http://localhost:8080/` in browser

5. And enjoy it.


Technical Details:
- Version: 2.0 (18.04.2025 release)
- v1: Using MySQL and OpenServer to save the notes
- v2: Using PostgreSQL with Docker to save the notes 
- v2: added simple tests 
- v2: added graceful shutdown
- There is no note editing feature now. There are also many more things that I would like to add or change in my program. I hope that over time I will improve my project.
- Creator: JustSMN
- Contact in telegram: @Just_Semen228
