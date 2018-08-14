SELECT
        people.Person.*,
        people.Prefix.prefixName,
        people.Adminpos.adminposName,
        people.Department.deptName
FROM
        people.Person
        Left Join people.Prefix ON people.Person.prefixId = people.Prefix.prefixId
        Left Join people.Adminpos ON people.Person.adminposId = people.Adminpos.adminposId
        Left Join people.Department ON people.Person.deptId = people.Department.deptId
WHERE
        Person.personCode LIKE '3%' AND people.Person.personId > 0
ORDER BY
        people.Person.adminId ASC,
        people.Person.levelposId DESC,
        people.Person.salary DESC