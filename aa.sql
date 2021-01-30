


select
 (user_id)) as pv
,post_id
from visa
group by
post_id
order by pv desc limit 10
