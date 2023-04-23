CREATE TABLE IF NOT EXISTS eletricity_bills(
    id INT PRIMARY KEY,
    service_id uuid NOT NULL,
    year INT NOT NULL,
    month INT NOT NULL,
  
    peak_kWh INT NOT NULL,
    peak_unit INT NOT NULL,
    peak_total INT NOT NULL,

    offpeak_kWh INT NOT NULL,
    offpeak_unit INT NOT NULL,
    offpeak_total INT NOT NULL,

    total INT NOT NULL
);

CREATE TABLE IF NOT EXISTS eletricity_bills_items(
    id INT PRIMARY KEY,
    bill_id INT REFERENCES eletricity_bills(id),
    label VARCHAR(100) NOT NULL,
    cost INT NOT NULL
);