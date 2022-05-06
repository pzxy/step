

insert into sh_dsj_zhywcx.t_ads_black_warn_collect partition(day='20220419')
select x.type,case when x.warncount is null then '0' else x.warncount end warncount,case when y.bkcount is null then '0' else y.bkcount end bkcount,current_date from
    (
        select t3.type,count(distinct(callernm)) warncount from
            (select t2.warnType type,t1.callerNum from
                (select callerNum from sh_dsj_ods.t_ods_bh_hmd_topic_incre_h where day>=date_format(date_add(current_date,-7),'yyyyMMdd') and opType='103') t1
                    left join
                (select callNumber,warnType from sh_dsj_ods.t_ods_xf_warn_topic_incre_h where day>=date_format(date_add(current_date,-7),'yyyyMMdd')) t2
                on t1.callerNum=t2.callNumber) t3
                left join
            (select callernm from sh_dsj_ods.t_ods_call_cdr_incre_h where day = '20220419' and ccsflag='206') t4
            on t3.callerNum=t4.callernm
        group by t3.type) x
        left join
    (
        select b.warn_type,count(distinct(callernm)) bkcount
        from
            (select callerNum from sh_dsj_ods.t_ods_bh_hmd_topic_incre_h where opType = '103' and day >= date_format(date_add(current_date,-7),'yyyyMMdd')) a
                left join
            (select callernm,warn_type from sh_dsj_dw.t_dwd_call_info_xd_incre_d where day = '20220419') b
            on a.callernm = b.callernm
        group by b.warn_type
    ) y
    on x.type = y.warn_type


insert into sh_dsj_zhywcx.t_ads_black_warn_collect partition(day='20220419')

delete from sh_dsj_zhywcx.t_ads_black_warn_collect partition(day='20220419')