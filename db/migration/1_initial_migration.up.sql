CREATE TABLE status (
    id INTEGER PRIMARY KEY,
    service TEXT NOT NULL,
    endpoint TEXT NOT NULL,
    status TEXT NOT NULL
);

-- TODO: Check if primary key order affect performance
CREATE TABLE requiring_service (
   service TEXT NOT NULL,
   endpoint TEXT NOT NULL,
   PRIMARY KEY (service, endpoint)
);

-- TODO: Check if primary key order affect performance
CREATE TABLE required_service (
  endpoint TEXT NOT NULL,
  service TEXT NOT NULL,
  PRIMARY KEY (endpoint, service)
);
