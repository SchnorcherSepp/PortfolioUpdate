# PortfolioUpdate – Country Weightings from Finanzfluss to Portfolio Performance

## Purpose
This tool reads a Portfolio Performance export in JSON format, loads the country weightings for each ISIN from Finanzfluss, and overwrites the existing country shares in the export. A new JSON file is then generated, which can be imported back into Portfolio Performance.

## How it works
1. **Load input:** Read `Regions_(MSCI).json` export from Portfolio Performance.
2. **Query Finanzfluss:** Load the corresponding Finanzfluss page for each ISIN using Playwright and parse the visible text.
3. **Extract countries:** Read the "Countries" section from the page content and convert percentages to numbers.
4. **Apply mappings:** Differing names are standardized (e.g., *Hong Kong* → *Hong Kong*, *United Kingdom* → *Great Britain*). The list can be expanded if necessary.
5. **Set weights:** The country weights found replace the instrument's previous weights.
6. **Write export:** Output the result as `Regions_(MSCI)_(ex.A)-2.json` (for re-import into Portfolio Performance).
