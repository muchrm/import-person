SELECT
                                        `PersonT`.*,
                                        `pf1`.`prefixName` AS `o1prefixName`,
                                        `pf2`.`prefixName` AS `o2prefixName`
                                FROM
                                        `PersonT`
                                        Left Join `Prefix` AS `pf1` ON `PersonT`.`o1prefixId` = `pf1`.`prefixId`
                                        Left Join `Prefix` AS `pf2` ON `PersonT`.`o2prefixId` = `pf2`.`prefixId`
                                WHERE
                                        `PersonT`.`personId` = 379