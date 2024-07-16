select * from regions
select * from countries

-- join
select r.region_id, region_name,country_id,country_name 
from regions as r join countries as c
on r.region_id=c.region_id
where r.region_id=2
order by r.region_name

-- join 2
select e.employee_id,first_name||''||last_name as full_name,email,hire_date,salary,
d.department_id,department_name
from employees as e join departments d
on e.department_id=d.department_id

-- count employee each department
select d.department_id,department_name, count(e.employee_id)as total_employee
from departments as d join employees e
on e.department_id = d.department_id
group by d.department_id,department_name
order by total_employee desc

-- display highsalary each department
select d.department_id,department_name,employee_id,first_name||''||last_name as full_name,salary as max_salary
from departments as d join employees e
on e.department_id = d.department_id
--where salary = (select max(salary) from employees where department_id=d.department_id )
order by department_name

-- display second highsalary each department
select d.department_id,department_name,employee_id,first_name||''||last_name as full_name,salary as max_salary
from departments as d join employees e
on e.department_id = d.department_id
where salary =(select salary from 
(select employee_id,salary ,row_number() over(order by employee_id)
from employees where department_id=d.department_id
order by salary desc)t
where row_number=2) 
order by department_name


select * from 
(select employee_id,salary ,row_number() over(order by employee_id)
from employees where department_id=8
order by salary desc)t
where row_number=2

select * from employees e 
where salary < (select max(salary) from employees)
and department_id=10
order by employee_id

-- display minimusalary each department
select d.department_id,department_name,employee_id,first_name||''||last_name as full_name,salary as max_salary
from departments as d join employees e
on e.department_id = d.department_id
where salary = (select min(salary) from employees where department_id=d.department_id )
order by department_name

select max(salary) from employees where department_id=11

select * from employees where salary=
(select max(salary) from employees)

select sum(salary) from employees

-- sum salary each department
select d.department_id,department_name,
sum(salary) as total_salary
from departments as d join employees e
on e.department_id = d.department_id
group by d.department_id,department_name
order by department_name

--- having sum
select d.department_id,department_name,
sum(salary) as total_salary
from departments as d join employees e
on e.department_id = d.department_id
group by d.department_id,department_name
having sum(salary) >=50000
order by department_name
