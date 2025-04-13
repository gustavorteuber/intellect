# Intellect - Auction Data Processing Tool

**Intellect** is a powerful data processing tool designed for managing and processing property auction data. It uses advanced technologies such as Amazon Bedrock (Claude 3) for data cleaning and transformation. This tool integrates seamlessly with MySQL for storing auction data. Developed and maintained by **Leiloes do Brasil**.

## Features

- **CSV Processing**: Reads auction property data from CSV files.
- **Data Transformation**: Cleans and transforms data using Amazon Bedrock (Claude 3).
- **MySQL Integration**: Saves processed data to a MySQL database.
- **CLI Interface**: Simple and intuitive command-line interface for running the application.

## Prerequisites

Before using **Intellect**, ensure that you have the following tools and environments set up:

- **Go**: The project is written in Go (Golang). You can download it from the official website: [Go Downloads](https://golang.org/dl/)
- **MySQL**: You must have MySQL installed to store the processed data. Create a database called `auctions` before running the application.
- **AWS**: You need an AWS account to use Amazon Bedrock (Claude 3). Set up AWS credentials with appropriate access.
