# Testcard

Easily find Stripe test cards using your terminal.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/joosto/testcard.git
cd testcard
```

2. And install the command line tool.

```bash
make install
```

## Usage

Find a Stripe test card:

```bash
$ testcard GB
4000058260000005        tok_gb_debit    United Kingdom  GB
```

Find a Stripe test card and only display the card number:

```bash
$ testcard -s GB
4000058260000005
```

Find a Stripe test card and pipe it to your clipboard ✨ (macOS only):

```bash
$ testcard -s GB | pbcopy
4000058260000005
```