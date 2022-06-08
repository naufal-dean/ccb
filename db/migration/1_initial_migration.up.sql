CREATE TABLE status (
    rd_service TEXT NOT NULL,
    rd_endpoint TEXT NOT NULL,
    status TEXT NOT NULL,
    expiry INTEGER NOT NULL,
    PRIMARY KEY (rd_service, rd_endpoint)
);

CREATE TABLE requiring_service (
   rg_service TEXT NOT NULL,
   service TEXT NOT NULL,
   endpoint TEXT NOT NULL,
   PRIMARY KEY (rg_service, service, endpoint)
);

CREATE TABLE required_service (
  endpoint TEXT NOT NULL,
  rd_service TEXT NOT NULL,
  rd_endpoint TEXT NOT NULL,
  PRIMARY KEY (endpoint, rd_service, rd_endpoint)
);
