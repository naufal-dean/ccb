-- TODO: add rd_method column
CREATE TABLE status (
    rd_service TEXT NOT NULL,
    rd_endpoint TEXT NOT NULL,
    status TEXT NOT NULL,
    PRIMARY KEY (rd_service, rd_endpoint)
);

-- TODO: Check if primary key order affect performance
-- TODO: add method column
CREATE TABLE requiring_service (
   rg_service TEXT NOT NULL,
   endpoint TEXT NOT NULL,
   PRIMARY KEY (rg_service, endpoint)
);

-- TODO: Check if primary key order affect performance
-- TODO: add rd_method and method columns
CREATE TABLE required_service (
  endpoint TEXT NOT NULL,
  rd_service TEXT NOT NULL,
  rd_endpoint TEXT NOT NULL,
  PRIMARY KEY (endpoint, rd_service, rd_endpoint)
);
