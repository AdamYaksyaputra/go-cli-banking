# go-cli-banking
File-based CLI banking system built with Go featuring account management, fund transfers, deposits, and compound interest calculation.

This project implements core banking operations including account creation, balance transfers, deposits, and compounded interest accrual using plain text file storage.

---

## Features

- Create bank accounts
- Transfer funds between users
- Add savings deposits
- Compound deposit interest calculation
- Transaction history logging
- File-based persistence without database

---

## Tech Stack

- Go (Golang)
- File System Storage
- CLI (Command Line Interface)

---

## Project Structure

```text
blu-cli-banking-system/
│
├── main.go
├── go.mod
├── README.md
│
├── services/
│   ├── account_service.go
│   ├── transfer_service.go
│   ├── deposit_service.go
│   └── interest_service.go
│
├── utils/
│   └── file_helper.go
│
├── accounts/
├── deposits/
└── history/
```

---

## File Structure

### Account File

Location:

```text
accounts/<name>.txt
```

Example:

```text
1500.00
[2026-05-18 10:00] Account Created
[2026-05-18 10:10] Transfer In +500
```

---

### Deposit File

Location:

```text
deposits/<name>.txt
```

Format:

```text
<amount>|<timestamp>
```

Example:

```text
500.00|2026-05-18T10:00:00
```

---

## Interest Rule

Deposits earn:

- **1% compounded interest per minute**

Formula:

\[
A = P(1+r)^t
\]

Where:

- `P` = initial deposit
- `r` = interest rate (0.01)
- `t` = elapsed minutes

---

## Getting Started

### Clone Repository

```bash
git clone https://github.com/your-username/blu-cli-banking-system.git
```

### Enter Project Directory

```bash
cd blu-cli-banking-system
```

### Run Application

```bash
go run main.go
```

---

## Available Commands

### 1. Create Account

```bash
go run main.go create_account Alice 1000
```

Expected Output:

```text
Account created successfully
```

---

### 2. Transfer Funds

```bash
go run main.go transfer Alice Bob 500
```

Expected Output:

```text
Transfer completed successfully
```

---

### 3. Add Deposit

```bash
go run main.go add_deposit Alice 500
```

Expected Output:

```text
Deposit added successfully
```

---

### 4. Accrue Interest

```bash
go run main.go accrue_interest
```

Expected Output:

```text
Interest accrued successfully
```

---

## Example Workflow

```bash
go run main.go create_account Alice 1000

go run main.go create_account Rudolph 2000

go run main.go add_deposit Alice 500

go run main.go transfer Rudolph Alice 500

go run main.go accrue_interest
```

---

## Validation Rules

- Account names must be unique
- Transfer balance must be sufficient
- Amount must be greater than zero
- Deposits do not reduce main balance
- Interest only affects deposits

## Author

Developed as part of a backend engineering technical assignment using Go.
