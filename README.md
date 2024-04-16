Based on the provided `main.go` file, here is a suggested README.md:

# SSL Certificate Expiry Checker

This Go program checks the SSL certificate expiry dates for a list of domains.

## Dependencies

- `crypto/tls`: For establishing TLS connections.
- `fmt`: For formatting output.
- `os`: For accessing system arguments.
- `strings`: For string manipulation.
- `time`: For time-related operations.
- `github.com/jedib0t/go-pretty/v6/table`: For creating a pretty table in the terminal.

## Usage

Run the program with a comma-separated list of domains as an argument:

```bash
go run main.go domain1,domain2,...
```

The program will establish a TLS connection to each domain on port 443, retrieve the SSL certificate, and calculate the number of days until the certificate expires. The results are displayed in a table sorted by the expiry date.

## Output

The output table has the following columns:

- `Domain`: The domain name.
- `Expires in (days)`: The number of days until the SSL certificate expires.
- `Expiration Date`: The exact date and time when the SSL certificate expires.

If a TLS connection cannot be established to a domain, the `Expires in (days)` and `Expiration Date` columns will show `-1` and `--` respectively.