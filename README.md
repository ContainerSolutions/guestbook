Guestbook
---

Goal: to write a stand-alone, single binary guestbook application in Go with PostgreSQL
as its backend storage.

# Roadmap
  - [x] Initial web app. No tests. Serves index page with guestbook form.
  - [x] SQL script for initialising tables in db
  - [ ] Flag in go app to run SQL script (`./guestbook -run-migrations`)
  - [ ] Flag to run web server
