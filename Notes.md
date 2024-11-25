
# My Notes

## Dependencies
- `Go v1.23.3` was used
- `Sqlc v1.27.0` was used

## App Notes
- Used the Hexagon pattern/structure
- Used middlewares to get app metrics(didn't add that of db for now), recover panic.
- Also used a Rate limiter to gaurd against such attacks
- Left minimal comments, but ones when i felt necesary
- Refactored make to reflect any change
- Used julienschmidt/httprouter for routing, which is easily swapable
- App and Db runs on same ports as before

## Testing Strategy
- Kindly find attached a .md file called testing_strategy

## Performance improvements
- Pagination of lists response

## Security 
- DbURl exposed in MakeFile
- Figured that since it's a test and you would run it out of the box

## SQLC
- also using this widely in a long time,i usual writing that layer from scratch.

## Feedback
- Please do provide feedback anyhow this goes, would be appreciated.
