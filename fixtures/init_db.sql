-- SQL for init
-- Create product table
CREATE TABLE IF NOT EXISTS product (
    productid VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Optional: Insert some sample data
INSERT INTO public.product ("productid", "productlevel", parentid, parentlevel, longdescription, shortdescription, imageurl, "productclass") 
VALUES ('PRD-sample123', 'PRD', 'LOREWI', 'PH1', 'Go sample, consectetur adipiscing elit. Suspendisse poten.', 'desc13279101', 'ttp://www.example.com/stracciatella.png', 'CONSUMABLE');