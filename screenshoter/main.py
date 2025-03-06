#!/usr/bin/env python3
import os
import asyncio
import argparse
from urllib.parse import urlparse
from playwright.async_api import async_playwright

async def take_screenshots(urls, output_dir):
    """Take screenshots of the given URLs using a single browser instance with multiple tabs."""
    async with async_playwright() as p:
        browser = await p.chromium.launch()
        
        # Create tasks for each URL
        tasks = []
        for url in urls:
            tasks.append(process_url(browser, url, output_dir))
        
        # Run all tasks concurrently
        await asyncio.gather(*tasks)
        
        # Close the browser when done
        await browser.close()

async def process_url(browser, url, output_dir):
    """Process a single URL in a new tab."""
    # Create a new page (tab) for this URL
    page = await browser.new_page()
    
    try:
        await page.goto(url, wait_until="networkidle", timeout=60000)
        
        # Get the page title for the filename
        title = await page.title()
        # Clean the title to make it a valid filename
        filename = "".join(c if c.isalnum() or c in " -_" else "_" for c in title).strip()
        if not filename:
            # Fallback to domain name if title is empty
            parsed_url = urlparse(url)
            filename = parsed_url.netloc.replace(':', '_')
            
        # Add timestamp to avoid overwriting
        import time
        timestamp = int(time.time())
        filename = f"{filename}_{timestamp}.png"
        
        # Ensure the output directory exists
        os.makedirs(output_dir, exist_ok=True)
        
        # Take the screenshot
        screenshot_path = os.path.join(output_dir, filename)
        await page.screenshot(path=screenshot_path, full_page=True)
        print(f"Screenshot saved: {screenshot_path}")
        
    except Exception as e:
        print(f"Error capturing {url}: {str(e)}")
    finally:
        await page.close()  # Close the tab when done

async def main():
    parser = argparse.ArgumentParser(description="Take screenshots of a list of URLs")
    parser.add_argument("--urls", nargs="+", help="List of URLs to screenshot")
    parser.add_argument("--file", help="File containing URLs (one per line)")
    parser.add_argument("--output", default="screenshots", help="Output directory for screenshots")
    
    args = parser.parse_args()
    
    urls = []
    if args.urls:
        urls.extend(args.urls)
    
    if args.file:
        try:
            with open(args.file, 'r') as f:
                file_urls = [line.strip() for line in f if line.strip()]
                urls.extend(file_urls)
        except Exception as e:
            print(f"Error reading URL file: {str(e)}")
    
    if not urls:
        print("No URLs provided. Use --urls or --file to specify URLs.")
        return
    
    # Process all URLs with a single browser instance
    await take_screenshots(urls, args.output)
    
    print(f"Completed. Screenshots saved to {os.path.abspath(args.output)}")

if __name__ == "__main__":
    asyncio.run(main())
