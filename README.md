# jurisdiction-tax-cacluator
A tax jurisdiction calculator built in Go. Enter your income, and it calculates
what you'd owe under the tax codes of several low-tax counties, so you can 
compare your actual take-home pay across jurisdictions. This project has the long-term 
goal of helping remote workers and high earners identify countries where their tax
burden and overall setup aligns with their goals.

## Currently supported
- Singapore
- UAE
- Bulgaria
- USA

## How to run
```bash
go run main.go
```
Enter your income in USD when prompted.

## Roadmap
- Add more countries
- Generalize the tax bracket logic into a reusable structure
- Possibly a web frontend
