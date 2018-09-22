SELECT 
  ROUND(
    6371 * ACOS(
      SIN(-6.597147 * PI() / 180) * SIN(restaurant_latitude * PI() / 180) + COS(-6.597147 * PI() / 180) * COS(restaurant_latitude * PI() / 180) * COS(
        (restaurant_longitude * PI() / 180) - (106.806038 * PI() / 180)
      )
    ),
    1
  ) AS distance,
  restaurant_name,
  restaurant_location,
  restaurant_latitude,
  restaurant_longitude 
FROM
  restaurant 
ORDER BY distance ASC 
