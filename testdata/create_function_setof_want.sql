CREATE OR REPLACE FUNCTION test_scalability(sql_txt int2) RETURNS SETOF record AS
'THIS IS UNPARSED BY THE INITIAL PARSER' LANGUAGE plpgsql IMMUTABLE STRICT;
