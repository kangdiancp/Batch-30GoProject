--subquery

select * 
from regions r join  countries c on r.region_id=c.region_id 
join locations l on c.country_id=l.country_id
where r.region_id=2


select * 
from employees e join departments d
on e.department_id=d.department_id
where d.location_id in (
select l.location_id
from regions r join  countries c on r.region_id=c.region_id 
join locations l on c.country_id=l.country_id
where r.region_id=1
)

-- delete
select * from regions where region_id in (5,6)

delete from regions where region_id in (5,6)