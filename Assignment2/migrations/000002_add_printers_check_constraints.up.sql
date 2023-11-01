ALTER TABLE printers ADD CONSTRAINT printers_battery_check CHECK (battery_left >= -1);
ALTER TABLE printers ADD CONSTRAINT supported_paper_sizes_length_check CHECK (array_length(supported_paper_sizes, 1) > 0);