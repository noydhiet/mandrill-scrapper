# SRG Radar Project

A web scraping application to identify potential manufacturers for pharmaceutical products based on comprehensive selection criteria.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Scraping Sources](#scraping-sources)
- [Selection Criteria](#selection-criteria)
- [Flow Diagram](#flow-diagram)
- [License](#license)

## Overview
This application helps pharmaceutical companies identify potential manufacturers for Active Pharmaceutical Ingredients (APIs) by:
1. Scraping data from multiple authoritative sources
2. Applying a comprehensive set of selection criteria
3. Presenting results in a structured format

## Features
- Multi-source data collection
- Comprehensive filtering system
- Priority-based ranking of manufacturers
- Detailed product information
- Regulatory compliance checks

## Installation
```bash
# Clone the repository
git clone https://github.com/noydhiet/mandrill-scrapper.git

# Install dependencies
go mod download
```

## Usage

Run the application in local environment:
```bash
# Run the application
go run main.go api
```

Test the application:
```bash
curl http://localhost:8080/v1/search?query=paracetamol
```

Run Worker in local environment
```bash
go run main.go worker
```

Support scraping worker
| Module | Frequency |
| ------ | --------- |
| Patent | Every day |


## Scraping Sources
### Patent Expiry
- [Pharsight](https://pharsight.greyb.com/drug-patent-expiration-lists)
- [Elixir](https://elixirdemo.greyb.com/drug-screener)

### Lawsuits
- [Motley Rice](https://www.motleyrice.com/medical-drugs)
- [Drugwatch](https://www.drugwatch.com/legal/)

### Product Recalls
- [FDA Recalls](https://www.fda.gov/safety/recalls-market-withdrawals-safety-alerts)
- [CPSC Recalls](https://www.cpsc.gov/Recalls)
- [Malaysia NPRA](https://www.npra.gov.my/index.php/en/consumers/safety-information/product-recall.html)
- [Singapore HSA](https://www.hsa.gov.sg/announcements)

## Selection Criteria
The system applies the following priority-based filtering:

1. **Registration Status**
   - Number of registered products in Malaysia, Singapore, and Indonesia (Max 5)
   
2. **Legal Compliance**
   - No active lawsuits related to adverse effects
   - No product recalls related to adverse effects

3. **Patent Status**
   - Prioritized by expiration timeline (soonest first)

4. **Manufacturer Details**
   - GMP certification status
   - Quality assurance certifications
   - Research and development capabilities
   - Therapeutic specialization

5. **Product Details**
   - Dosage forms
   - Storage requirements
   - Usage frequency
   - Pricing information

## Flow Diagram
1. API Filtering
   - Number of registered products → Lawsuits → Product recalls → Patent status

2. Manufacturer Filtering
   - List of manufacturers → Distributors → GMP → Quality → R&D → Specialization → Pipeline

3. Product Filtering
   - Product list → Storage → Usage → Pricing

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
