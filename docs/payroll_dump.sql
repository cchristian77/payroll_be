PGDMP                         }            payroll    14.7 (Debian 14.7-1.pgdg110+1)    14.18 W    q           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            r           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            s           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            t           1262    16384    payroll    DATABASE     [   CREATE DATABASE payroll WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';
    DROP DATABASE payroll;
                admin    false            ^           1247    16536    reimbursement_status    TYPE     O   CREATE TYPE public.reimbursement_status AS ENUM (
    'PENDING',
    'PAID'
);
 '   DROP TYPE public.reimbursement_status;
       public          admin    false            I           1247    16394 	   user_role    TYPE     B   CREATE TYPE public.user_role AS ENUM (
    'ADMIN',
    'USER'
);
    DROP TYPE public.user_role;
       public          admin    false            �            1259    16483    attendances    TABLE     �  CREATE TABLE public.attendances (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    created_by bigint NOT NULL,
    updated_by bigint,
    user_id bigint,
    date date NOT NULL,
    check_in timestamp without time zone NOT NULL,
    check_out timestamp without time zone
);
    DROP TABLE public.attendances;
       public         heap    admin    false            �            1259    16482    attendances_id_seq    SEQUENCE     �   CREATE SEQUENCE public.attendances_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.attendances_id_seq;
       public          admin    false    220            u           0    0    attendances_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.attendances_id_seq OWNED BY public.attendances.id;
          public          admin    false    219            �            1259    16386    goose_db_version    TABLE     �   CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);
 $   DROP TABLE public.goose_db_version;
       public         heap    admin    false            �            1259    16385    goose_db_version_id_seq    SEQUENCE     �   CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.goose_db_version_id_seq;
       public          admin    false    210            v           0    0    goose_db_version_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;
          public          admin    false    209            �            1259    16508 	   overtimes    TABLE     f  CREATE TABLE public.overtimes (
    attendance_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by bigint NOT NULL,
    updated_by bigint,
    user_id bigint,
    date date NOT NULL,
    duration smallint NOT NULL
);
    DROP TABLE public.overtimes;
       public         heap    admin    false            �            1259    16432    payroll_periods    TABLE     �  CREATE TABLE public.payroll_periods (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by bigint NOT NULL,
    updated_by bigint,
    start_date date NOT NULL,
    end_date date NOT NULL,
    payroll_run_at timestamp without time zone
);
 #   DROP TABLE public.payroll_periods;
       public         heap    admin    false            �            1259    16431    payroll_periods_id_seq    SEQUENCE     �   CREATE SEQUENCE public.payroll_periods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.payroll_periods_id_seq;
       public          admin    false    216            w           0    0    payroll_periods_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.payroll_periods_id_seq OWNED BY public.payroll_periods.id;
          public          admin    false    215            �            1259    16453    payslips    TABLE     �  CREATE TABLE public.payslips (
    id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by bigint NOT NULL,
    updated_by bigint,
    user_id bigint NOT NULL,
    payroll_period_id bigint NOT NULL,
    total_attendance_days integer NOT NULL,
    total_overtime_days integer NOT NULL,
    total_overtime_hours integer NOT NULL,
    total_reimbursements bigint NOT NULL,
    base_salary bigint NOT NULL,
    attendance_pay bigint NOT NULL,
    overtime_pay bigint NOT NULL,
    reimbursement_pay bigint NOT NULL,
    total_salary bigint NOT NULL
);
    DROP TABLE public.payslips;
       public         heap    admin    false            �            1259    16452    payslips_id_seq    SEQUENCE     �   ALTER TABLE public.payslips ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.payslips_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          admin    false    218            �            1259    16542    reimbursements    TABLE     �  CREATE TABLE public.reimbursements (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by bigint NOT NULL,
    updated_by bigint,
    user_id bigint NOT NULL,
    description text,
    amount bigint NOT NULL,
    status public.reimbursement_status NOT NULL,
    payslip_id bigint,
    reimbursed_at timestamp without time zone
);
 "   DROP TABLE public.reimbursements;
       public         heap    admin    false    862            �            1259    16541    reimbursements_id_seq    SEQUENCE     �   CREATE SEQUENCE public.reimbursements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.reimbursements_id_seq;
       public          admin    false    223            x           0    0    reimbursements_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.reimbursements_id_seq OWNED BY public.reimbursements.id;
          public          admin    false    222            �            1259    16573    request_logs    TABLE     �  CREATE TABLE public.request_logs (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    request_id character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    activity character varying(255) NOT NULL,
    entity character varying(255) NOT NULL,
    reference_id bigint NOT NULL,
    client_ip character varying(255) NOT NULL
);
     DROP TABLE public.request_logs;
       public         heap    admin    false            �            1259    16572    request_logs_id_seq    SEQUENCE     �   CREATE SEQUENCE public.request_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.request_logs_id_seq;
       public          admin    false    225            y           0    0    request_logs_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.request_logs_id_seq OWNED BY public.request_logs.id;
          public          admin    false    224            �            1259    16413    sessions    TABLE     �  CREATE TABLE public.sessions (
    id integer NOT NULL,
    session_id uuid DEFAULT gen_random_uuid(),
    user_id integer NOT NULL,
    access_token text NOT NULL,
    access_token_expires_at timestamp with time zone NOT NULL,
    access_token_created_at timestamp with time zone DEFAULT now() NOT NULL,
    user_agent character varying(255) NOT NULL,
    client_ip character varying(255) NOT NULL
);
    DROP TABLE public.sessions;
       public         heap    admin    false            �            1259    16412    sessions_id_seq    SEQUENCE     �   CREATE SEQUENCE public.sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.sessions_id_seq;
       public          admin    false    214            z           0    0    sessions_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;
          public          admin    false    213            �            1259    16400    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    username character varying(100) NOT NULL,
    password text NOT NULL,
    full_name character varying(255),
    role public.user_role NOT NULL,
    base_salary bigint NOT NULL
);
    DROP TABLE public.users;
       public         heap    admin    false    841            �            1259    16399    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          admin    false    212            {           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          admin    false    211            �           2604    16486    attendances id    DEFAULT     p   ALTER TABLE ONLY public.attendances ALTER COLUMN id SET DEFAULT nextval('public.attendances_id_seq'::regclass);
 =   ALTER TABLE public.attendances ALTER COLUMN id DROP DEFAULT;
       public          admin    false    220    219    220            �           2604    16389    goose_db_version id    DEFAULT     z   ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);
 B   ALTER TABLE public.goose_db_version ALTER COLUMN id DROP DEFAULT;
       public          admin    false    209    210    210            �           2604    16435    payroll_periods id    DEFAULT     x   ALTER TABLE ONLY public.payroll_periods ALTER COLUMN id SET DEFAULT nextval('public.payroll_periods_id_seq'::regclass);
 A   ALTER TABLE public.payroll_periods ALTER COLUMN id DROP DEFAULT;
       public          admin    false    215    216    216            �           2604    16545    reimbursements id    DEFAULT     v   ALTER TABLE ONLY public.reimbursements ALTER COLUMN id SET DEFAULT nextval('public.reimbursements_id_seq'::regclass);
 @   ALTER TABLE public.reimbursements ALTER COLUMN id DROP DEFAULT;
       public          admin    false    223    222    223            �           2604    16576    request_logs id    DEFAULT     r   ALTER TABLE ONLY public.request_logs ALTER COLUMN id SET DEFAULT nextval('public.request_logs_id_seq'::regclass);
 >   ALTER TABLE public.request_logs ALTER COLUMN id DROP DEFAULT;
       public          admin    false    224    225    225            �           2604    16416    sessions id    DEFAULT     j   ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);
 :   ALTER TABLE public.sessions ALTER COLUMN id DROP DEFAULT;
       public          admin    false    214    213    214            �           2604    16403    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          admin    false    211    212    212            i          0    16483    attendances 
   TABLE DATA           �   COPY public.attendances (id, created_at, updated_at, deleted_at, created_by, updated_by, user_id, date, check_in, check_out) FROM stdin;
    public          admin    false    220   �t       _          0    16386    goose_db_version 
   TABLE DATA           N   COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
    public          admin    false    210   7u       j          0    16508 	   overtimes 
   TABLE DATA           {   COPY public.overtimes (attendance_id, created_at, updated_at, created_by, updated_by, user_id, date, duration) FROM stdin;
    public          admin    false    221   �u       e          0    16432    payroll_periods 
   TABLE DATA           �   COPY public.payroll_periods (id, created_at, updated_at, created_by, updated_by, start_date, end_date, payroll_run_at) FROM stdin;
    public          admin    false    216   v       g          0    16453    payslips 
   TABLE DATA             COPY public.payslips (id, created_at, updated_at, created_by, updated_by, user_id, payroll_period_id, total_attendance_days, total_overtime_days, total_overtime_hours, total_reimbursements, base_salary, attendance_pay, overtime_pay, reimbursement_pay, total_salary) FROM stdin;
    public          admin    false    218   �v       l          0    16542    reimbursements 
   TABLE DATA           �   COPY public.reimbursements (id, created_at, updated_at, created_by, updated_by, user_id, description, amount, status, payslip_id, reimbursed_at) FROM stdin;
    public          admin    false    223   �{       n          0    16573    request_logs 
   TABLE DATA           �   COPY public.request_logs (id, created_at, updated_at, request_id, user_id, activity, entity, reference_id, client_ip) FROM stdin;
    public          admin    false    225   |       c          0    16413    sessions 
   TABLE DATA           �   COPY public.sessions (id, session_id, user_id, access_token, access_token_expires_at, access_token_created_at, user_agent, client_ip) FROM stdin;
    public          admin    false    214    |       a          0    16400    users 
   TABLE DATA           y   COPY public.users (id, created_at, updated_at, deleted_at, username, password, full_name, role, base_salary) FROM stdin;
    public          admin    false    212   �       |           0    0    attendances_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.attendances_id_seq', 1, true);
          public          admin    false    219            }           0    0    goose_db_version_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.goose_db_version_id_seq', 8, true);
          public          admin    false    209            ~           0    0    payroll_periods_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.payroll_periods_id_seq', 7, true);
          public          admin    false    215                       0    0    payslips_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.payslips_id_seq', 921, true);
          public          admin    false    217            �           0    0    reimbursements_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.reimbursements_id_seq', 2, true);
          public          admin    false    222            �           0    0    request_logs_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.request_logs_id_seq', 1, false);
          public          admin    false    224            �           0    0    sessions_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.sessions_id_seq', 15, true);
          public          admin    false    213            �           0    0    users_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.users_id_seq', 102, true);
          public          admin    false    211            �           2606    16490    attendances attendances_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_pkey;
       public            admin    false    220            �           2606    16492 (   attendances attendances_user_id_date_key 
   CONSTRAINT     l   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_user_id_date_key UNIQUE (user_id, date);
 R   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_user_id_date_key;
       public            admin    false    220    220            �           2606    16392 &   goose_db_version goose_db_version_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.goose_db_version DROP CONSTRAINT goose_db_version_pkey;
       public            admin    false    210            �           2606    16514    overtimes overtimes_pkey 
   CONSTRAINT     a   ALTER TABLE ONLY public.overtimes
    ADD CONSTRAINT overtimes_pkey PRIMARY KEY (attendance_id);
 B   ALTER TABLE ONLY public.overtimes DROP CONSTRAINT overtimes_pkey;
       public            admin    false    221            �           2606    16439 $   payroll_periods payroll_periods_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.payroll_periods
    ADD CONSTRAINT payroll_periods_pkey PRIMARY KEY (id);
 N   ALTER TABLE ONLY public.payroll_periods DROP CONSTRAINT payroll_periods_pkey;
       public            admin    false    216            �           2606    16441 7   payroll_periods payroll_periods_start_date_end_date_key 
   CONSTRAINT     �   ALTER TABLE ONLY public.payroll_periods
    ADD CONSTRAINT payroll_periods_start_date_end_date_key UNIQUE (start_date, end_date);
 a   ALTER TABLE ONLY public.payroll_periods DROP CONSTRAINT payroll_periods_start_date_end_date_key;
       public            admin    false    216    216            �           2606    16459    payslips payslips_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_pkey;
       public            admin    false    218            �           2606    16461 /   payslips payslips_user_id_payroll_period_id_key 
   CONSTRAINT     �   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_user_id_payroll_period_id_key UNIQUE (user_id, payroll_period_id);
 Y   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_user_id_payroll_period_id_key;
       public            admin    false    218    218            �           2606    16551 "   reimbursements reimbursements_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.reimbursements
    ADD CONSTRAINT reimbursements_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.reimbursements DROP CONSTRAINT reimbursements_pkey;
       public            admin    false    223            �           2606    16582    request_logs request_logs_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.request_logs
    ADD CONSTRAINT request_logs_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.request_logs DROP CONSTRAINT request_logs_pkey;
       public            admin    false    225            �           2606    16423    sessions sessions_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_pkey;
       public            admin    false    214            �           2606    16425     sessions sessions_session_id_key 
   CONSTRAINT     a   ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_session_id_key UNIQUE (session_id);
 J   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_session_id_key;
       public            admin    false    214            �           2606    16409    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            admin    false    212            �           2606    16411    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            admin    false    212            �           2606    16493 '   attendances attendances_created_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);
 Q   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_created_by_fkey;
       public          admin    false    220    212    3239            �           2606    16498 '   attendances attendances_updated_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES public.users(id);
 Q   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_updated_by_fkey;
       public          admin    false    212    220    3239            �           2606    16503 $   attendances attendances_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 N   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_user_id_fkey;
       public          admin    false    220    212    3239            �           2606    16515 &   overtimes overtimes_attendance_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.overtimes
    ADD CONSTRAINT overtimes_attendance_id_fkey FOREIGN KEY (attendance_id) REFERENCES public.attendances(id);
 P   ALTER TABLE ONLY public.overtimes DROP CONSTRAINT overtimes_attendance_id_fkey;
       public          admin    false    3255    221    220            �           2606    16520 #   overtimes overtimes_created_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.overtimes
    ADD CONSTRAINT overtimes_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);
 M   ALTER TABLE ONLY public.overtimes DROP CONSTRAINT overtimes_created_by_fkey;
       public          admin    false    212    3239    221            �           2606    16525 #   overtimes overtimes_updated_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.overtimes
    ADD CONSTRAINT overtimes_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES public.users(id);
 M   ALTER TABLE ONLY public.overtimes DROP CONSTRAINT overtimes_updated_by_fkey;
       public          admin    false    221    212    3239            �           2606    16530     overtimes overtimes_user_id_fkey    FK CONSTRAINT        ALTER TABLE ONLY public.overtimes
    ADD CONSTRAINT overtimes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 J   ALTER TABLE ONLY public.overtimes DROP CONSTRAINT overtimes_user_id_fkey;
       public          admin    false    221    3239    212            �           2606    16442 /   payroll_periods payroll_periods_created_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.payroll_periods
    ADD CONSTRAINT payroll_periods_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);
 Y   ALTER TABLE ONLY public.payroll_periods DROP CONSTRAINT payroll_periods_created_by_fkey;
       public          admin    false    212    216    3239            �           2606    16447 /   payroll_periods payroll_periods_updated_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.payroll_periods
    ADD CONSTRAINT payroll_periods_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES public.users(id);
 Y   ALTER TABLE ONLY public.payroll_periods DROP CONSTRAINT payroll_periods_updated_by_fkey;
       public          admin    false    3239    212    216            �           2606    16462 !   payslips payslips_created_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);
 K   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_created_by_fkey;
       public          admin    false    212    3239    218            �           2606    16477 (   payslips payslips_payroll_period_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_payroll_period_id_fkey FOREIGN KEY (payroll_period_id) REFERENCES public.payroll_periods(id);
 R   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_payroll_period_id_fkey;
       public          admin    false    218    3247    216            �           2606    16467 !   payslips payslips_updated_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES public.users(id);
 K   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_updated_by_fkey;
       public          admin    false    3239    212    218            �           2606    16472    payslips payslips_user_id_fkey    FK CONSTRAINT     }   ALTER TABLE ONLY public.payslips
    ADD CONSTRAINT payslips_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 H   ALTER TABLE ONLY public.payslips DROP CONSTRAINT payslips_user_id_fkey;
       public          admin    false    212    218    3239            �           2606    16552 -   reimbursements reimbursements_created_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.reimbursements
    ADD CONSTRAINT reimbursements_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);
 W   ALTER TABLE ONLY public.reimbursements DROP CONSTRAINT reimbursements_created_by_fkey;
       public          admin    false    3239    223    212            �           2606    16567 -   reimbursements reimbursements_payslip_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.reimbursements
    ADD CONSTRAINT reimbursements_payslip_id_fkey FOREIGN KEY (payslip_id) REFERENCES public.payslips(id);
 W   ALTER TABLE ONLY public.reimbursements DROP CONSTRAINT reimbursements_payslip_id_fkey;
       public          admin    false    218    3251    223            �           2606    16557 -   reimbursements reimbursements_updated_by_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.reimbursements
    ADD CONSTRAINT reimbursements_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES public.users(id);
 W   ALTER TABLE ONLY public.reimbursements DROP CONSTRAINT reimbursements_updated_by_fkey;
       public          admin    false    3239    223    212            �           2606    16562 *   reimbursements reimbursements_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.reimbursements
    ADD CONSTRAINT reimbursements_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 T   ALTER TABLE ONLY public.reimbursements DROP CONSTRAINT reimbursements_user_id_fkey;
       public          admin    false    212    3239    223            �           2606    16583 &   request_logs request_logs_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.request_logs
    ADD CONSTRAINT request_logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 P   ALTER TABLE ONLY public.request_logs DROP CONSTRAINT request_logs_user_id_fkey;
       public          admin    false    3239    225    212            �           2606    16426    sessions sessions_user_id_fkey    FK CONSTRAINT     }   ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 H   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_user_id_fkey;
       public          admin    false    214    212    3239            i   G   x�3�4202�50�50W02�24�24�31��03ǔ2��322421���440�c�J��3�421����� G�      _   z   x�uϻAE�x���
��Z\���G,Dȑ����Y`��qq�D���*�&��g�cЦ. iӮ�IWY׆��!V��؄m����)�sG�_�96�`�U���);T\��7}�_C}      j   6   x�3�4202�50�50W02�24�20�3��00�#ch`����8��b���� �M�      e   p   x�}��	�@��M6�!?;I�"�����C�mȗ�`b114�&��� �6��ز��l�wt9 廜C�{v�S���߫���0cudn3�������֪��i�(-�LD_��6�      g   �  x�}��q1D�RN�.� �8G���05@�|�n�bU��.�9��1ڈ�m�l�G��1��k��*�m|�����O��h�?���G'�0**��pF�g�$��	�M�ORh�h'�y��
N��%��O	��}���K��Ρ')����ؽ2��H8�i��$��i�nDν�3��.�4�8mv�.�d�� 31#Х�FF���Ȑd�TԈ���']:j�ͷ6�J���ڻ�v���XE�KK�!6��@��ڻ��_7Y�+]�j1��~%؉����闌@���mXs2 �.=�wOm�9��2]z��g-�.=�wOoaAD�t��"��=#Х��p�M�$#Х����9���T'Ƹٮ+]��d��nA�LF(줫A\���<{F�KW)5+#Х�A&FL:g2]��:w�`�XF�KW��:��"[�J

��u2U3]�ʖ�[�/֨']�����$�@��i��v�T�tij�)�N;샟�$i�$����	@�����nw���ti)9!���٢����IT<=؆���R6�O�od��t�\Ζ��@�������;�+]z:ߛ���yﴊ@���w]���Z���--���{���p� -]�z�M���.M]�z����ti�"�����.M%��џti*}�{j!S&#Х����L*��@��.��d��f�.M]�)B*]����wޱɚ��U�P�/r�R�狼����{�.]ݤ�1���V�t��1�����M\�{�g�z"Х�LwzY���UMi*���F
��@����d�A�:�.Me����v����ML]�l��ti�&#�V#3�Ip.==��u薕��SR���`KgF�KOYIqK�N&SF�KOYI��$���.=e��g=���Tr��w�a�[F�KSY�x��R!#Х��mrT��TrX�N�LE�KS�5f�m�LS�.)WO{w5���SE�+W�|���JP���CV��꓀�L=��(�`�uE�+SOc��dٯte*{cwR�T�2�42��;�Юte�!�`��1��L=�t��I�W�2��j���g�W���NLu;�R�"Х�������t�tij'Cw��L�'[��I�c�eU�4��盃���.M�����#*=e�7��{��@��2]nqN��@���kJ$]����N<]=����o=HQ�w�bݚ�<|F�xi*����'+D2�Ə�kO����z��ˏ��?_{A4dw̢}��������o�      l   s   x���=
�@��z��@�7�c��I<@@����"�)^���H�Q{�:�0pv5Q�H���\|01d_���߮m#/ �<���!�E�����O^���!w�)�]�2�      n      x������ � �      c   �  x����n���������S����q�11��#E�!	��7�R�t -�(RmA}^Y���1�t�!#�:�iJP,��X�$��Xd�i';ZE<JJ��ƳvL�r�2�LE��q����-���1<*O8Ռ�fQ��U���1��c��k�W�5��i�"4W-�����ʣ�W��`��}.�������׬�v�z�?_��n�9��o�mr�߽���z�G��`�{��!IG�2�~�o���Ȼ�4$�>����f/�5Z^�>of����PL�a�;�o�a�?�1f���ˮ�6��ͮl��j?8��;77����4kGBC\����ZʣT��/ n�Z����+a/�e�� q�Us����?8:t@��t�.�cp%b�. �O�A,�Xt{ci�� @��D������L�N���ۻ�Q����"�_���KHiF���9��(����2j�zB���p�s@l��+��g���E��HqV	�ǝj��k��S�b_Q��� ~}HG�ۍ�8
-}%��c�VT�U��
�r<|&Kٝ��bJn���G��}�9b�I�L'��P���OqJ$�2�iY��WE@X��O�� 1�=����Z;�wڴa}]��Q�����£�qco���>���nwB�J<�n������h��wkUPk�}@̌L���Xt8c��Rq���Ck��In��E<����9;:�`��8`n��au���ppg1���#�S����p� uӪJa�~_F�!���� ���'L�>�>X��1=� �}t/�t9ݎ7�GU��C��Ma�+�βa��g�o��ݦG4}�{�ێ��|u���LW����$F&s#K�y�8�%t�n�8���D&	M�R��ۍ��4����W�W�SG8����p횽:hWmP�s]G@�g�b���� ;Գ���E.���;�9��]Lryo�;Ǚ����v��#&�"�_��kB�i �9�Dl ��EY�SÐk_����]xU���pj�S��M��U�vL]�+�u�L�'�"�ֻq7y����֟L���v�I6,�_V�[-voC����{����?G�w���1E8�@&0��I�9Ny��%���{4��G��U�� ��1(����� ���i��t}-b8���|�8�U�� �'F�tk[������q��[>!O�C�m3����1�7������s����TPEyB�q��qDe��,K�/�x^H��^��%�xxEH�(��ӠME���2�%d�W��e^b�r������ݴQܫ{w�`�����̾u�/�����	�hI2j0$Ʊ�t,9���N�(I���Ds�UU���uM1���������FS�	�r�����ٙ;�"��ʉY��Y��f7�l��[c��ʎ��Q�*O�oU�U�����������w,��aL;q�1�	C"���k���p�\K�!#��Z����N����p~�	s�OU1�wa_�`�������c�Lu�}�ؔbx�2od��n�*���޻s��B�V�x�6��OQf���u��A��)�YV�Efp�M��S�����3�j��uXuO9.�
�c�`�Pp�{�:�������u�S��b^���)w�<�͇�z�V%��m�,��\�3�5�7<
o��9���9���0�<���i�X��-ӑ�a��,Ch3��k-n�T('�)CC,��X�����'-�i�
iV���ϴ�D��ƹJ��o�����^�8w�gML}�'Q�a��i�69/�7D�p���>�^ёY�=��"
Җa��T�Lgz*�(�ŗ�-\�����p`���mw�W�I��jL�Ō����vE�U�݊��]���_Y��iel�׵1��/s�}yՇp�]�L&��[�ϙU�A}d��>�	�}�9�������?��      a      x���ǲ�J���OQ������E�'a$���{W������P[!�!�/�J2�_�����@���0�O������K�5�Y��;����,0��O���Ʈ�0��
!I1S �Մ��H����F�e�ܟ���˵y��������w"��ȇٔ\������B^C8S�a�"�8ڭl��)�G�ϵ/j'���=��&�F铹h1LC��ņ�r�f.H�Q�≣��b&/�<���!M�K�~�f��p���X&sӓP��7��/��N9��1a�BA���;��D��D�7�نm����w�C�P9n(eP�=���|�O�^	������k"�X"��(�8����z�^yC_�+��+q�X�Ћ��	2|F6��D�D�7Q#Fp��*܏���r<�W���D��ތ���[S5�������k"�X"��.��.W��u.p������ ԧ��L"Z}$朼.BmŌ����;��D��D�77:�(�Q��E2�;۵>Ӫչ���.���î=N鷆|!��F�5�D#�:�q` �TvcrM|���d�T�D߷b�ޞ���Jd\L��(�ߑ�]$���������f�����~�.2L�4�O���Fi�`kT|���K�oW�� �s>
�&��E��-0ZNN�S�7z��=�F�P0���|��u.��8�#|B���oxk�
!�,�~�/��P`������:�X�!��X$A��}|�8��>�7��-��ZZ��U����裨�y��aM�cKn��Ij�;��w���ׯWAzVO�c �^��<��,��0]

_�"]k0�U�Y����D�'0�U�ηW�45�hB���;x<�$s͆%�����'Ԏq����9m�`^Zi�t��P�:]�1bͩ�N�t�k��9O}��Ǘs���_j'9|�rx����"~hO�Iۏ�����g�7D��p��@(:8�X��Z_�/	;�ლ��L���iK�(�2��u��J�C�W�b}�j�HT�lɷV�)�����A��tH�����|��׀�cz�̬�S�S&{BueW���t�ש���AёM�W���Wv�5�5-�@�C���1�+���k8���w��Y��g�C���AёMt���� B�vYP�ix��VZPf,��7��[ét{(�	��j':r��DoW�!Dե5j���-�zx�����LJ����כ!�l+����Nt���&:�3��g��e�¡X���t>���k�	 �n!���[�SY���9(:��.%�S��疿Ө���w{hy��|��ޏ�ȣ��N�_�}�>ى��D��j�R��<E�́��3FF�s�k��!!?���@�8z~&dd':rPtd��
��._�J���)�+0��R��o���!���S.O�[�z�O���DG��l�����{�%Q)
�5�G�.�['r���	'�B+Arl��ß��4��DG��l�[IJ���'O=-�W�r���ڪ(�rNڒa�^��!�O����AёMt�jM�_i��͉��PG�Cɀ�^|�� *�)�)�얀�N�DG��n�('��߃7yw���9.��֫�����V���%ڨ�2���������A��MtOԐ�ԇDjR�MIx㭱���]z^�Z��@ ����cm���	�_^9(:���-��L�A�fINBȲv�|��R��K�=�H��J��N}ct':zt�e�<���P>"9F��}��VL��-��8��
�����=x�B۟�݉��D�WR�Ѭ�A�%�֐�*Z�)�+�ϰwV[��Te�"ƾ�?}���݉��D���7��P!O�|�D��X�y�G/�ހ�J��T����p�|C�DG��n�k���-�>B�џ}\3�p�a����l��Z&�L��݉��D�!��5��Y�n�W����9�þ��%Ǧ�	��+0�?W/�,%�;�у����ۙ�z-���ɳ��]C��s���D�(�N��ju��h��w�w��D���������\+�g�(y~{�wy�:�u� ����wO�V5��!�胙�F�݉��D�y�-�xI05|0�C$��o���Nٺ���#��_�˭k��.r�D���m������u�V��r�ff���/�N��B3�B�4�2¿�O���D���m�Sg�!�����J����Z����&�z�E��ؽ��ŵ�;!c�5�c���-����%�Y�}:=kZH{(��дT�0���2A�3�3�`;ѱ���踠j�Ϡ��	 J�%d�8�_rl@���|�AL�h�(��}�Nt���&�$��=��K�z�wt5����;y@^�Q��+`���
���������AѱM��m�lQF_'d�z�@�v
W��i���=�\c�3[��E�wL�D���m����J�Bb�D.�%�~+�(��w1���}�g���Hs�˹����D���m�_˳���$��81O��
� j5j���nb��6�CŜ�N�D���m�sZ�p��m�Ԣ�;e������-�50M]>��=�!�Qp��4�Nt����&�7��h����� |�;�>�Y�d�B�lD��G����)�t|':~Pt|����[A�"�;���]yl�3�7�J�S9��A4��2\Ŋ��>��?(:��Ϊ�� ��r�n�)�OK���,���TGq��p�,~ ����*��?(:���b0~'�^�<k�o��E�M]������&-{���
Pꋚ%������E�7��SEν���H����=�d�b!T�o�L��E�N��*6/�|��R;��C7��
|�#r3T*Z^y6�ȝ�b��}��Ÿ�ݰI�O9��D���o�gV��uU�$�P��\���5���_aƫA�\_�#'g���L�)����A��Mt�Ǝ�/�F��a�k	2�l#V��a��:��'h��yL��/˧m�w��E�7��B�m�q;��g&��7ώ�|Ɂ'C	Q��G�2c�4x��f�7�Nt���&:��Jx��#?G��Τ���%�D�"8Ϊ͝4����jY�C�D'�Nl��2�Z!��U-f7���K㉂O���H�Ή�����[*s3G?�ϣ5b':qPtb}��ҫ�f�5���H	��27�闲t��-�Ԡ�Jm!�o����AщMt����qÐ��x�c/�2�,iA�ȡNڙ���vΧO�t.<����AщM��ո����i�`Gr�2Y��/�z��"�P�x+���>�F��>}�b':qPtb=����1s��d�;�>��x+(�*�q-B��b? g��M�n���n9(:��~M	/a��Tұ�F��dKa�i�m�|�#6��N�����ӧq v�Gw�l��q�7�i�����C���;@�0����.�d�A�xt��w�}\K�D'�Nl�c��a�NSX��T�f5B�Tı�X���-֪\�yja��@߉N��D����Ħޑ�8s���N��\��o_���Xd�ャ�?�4`�ħ&v��E'6љ���b@�y�f��)3�(��;)�~�)�L[_Jk(3e�����A��M�%���%�Г��),��lF����i/*�E��1Ԃ�I=��
���Nt���&zPŀbA��rB���tC��.n�\�#23Z��93B���}�Nt���&�Dz	C�6�V&hH�Oa�@*7��Ab�\��AP(C	g�	����<(:��n����<z�M�*�>�y4�j%��+���5�@��SO�o����A��M��=�SI�'	�hj���u��&[�-Pu/`IM$����BLܾ���Nt���&�խ�I���-u��0d��KЌ��z��[��o����@�oY<(:���:#�W���B2������-A��/�4y�۴�&Ƅ�B�Ɓ܉Nݶ��>I'1Ω�N$	�7D�,y\e�΋Q �pQ<ԃ�N]ջwT~�"w��E'7�3mH�7 �  �G�
�JS��dpl�# dM|-�:"#^��5�+_}�ar':uPt�O����W�8ȃUx�	gb�a'�i�uL�;~��k\�J��
���S�Nt���&:ܺBIغ#z����qH�g�w=�Tn�z �-�c�0�ɕ>�0��:(:���E����FM�Mp�$��mm����EhVa0+^U����Nt��Ej':uPtj=hlݕ
!�&'��j�o�) �[x�C	C�L�H��Ƥ�O{�S�Nt���&�8�i�PCA�;���:��Vh�e�3f�b�ު��kq���ݝ��AѩM�%� bXG���0�2�q�T��j�\�����@�d5{��������AѩMt\Y��V���k�9�V\��ft��"Ê����f��I3���;�v�SE�6��k�H��3�YG�~3�}(*�'��]?�ó���W���>�'�߇~Ptj="/����"�����g:GG�uS,S�W_�N�Q�:�\�O�@�D���E�D���%�ƙ�1�~p�J��Ќ)��7���U�g&��|j':}Pt��utIGX���+��R��ӛ׽.���l��6����f�iP�o��?-��>(:��K�1':.�V�PB�q�r���.OJ"ѹ��3E<��z����g>k��N������P�藎p�*��GB` ��Y���5s�a\������tķν��'��>(:���ԩߜ���I�qY��@k<�R����Ji�T�ob��s��;�郢ӛ�2���c�u�]2\S�F�4\hA��;���2����xN�O=E�D��No�S�̀,Y�l�_����pK-]�]�,Q\їdr����'���)��>(:��K=�=�O=�4e'F�i������H��!���i��Oȯ�컎N�D��No���tz^'Q�r���Z��Ӯ�?���\��р*�V� _n7�o�G�D��No��i�ia���3���+R���ﶈ�����pqc��;�E[�������No��g٤�����W׊�>�^g���o�\�F֖�W�K`�= 8�v��ש�u�����t4a�w�ɜ��[���1!t�����8�����g��|3@����'��G���g�6�������{Z�g�K�e��_x��t�I�n��eΚ�s������G鳨�K(@�F���L��	�x�fآwe�y'��J�4�5V��� ��D�]4��_�J�/�w�?������)�^     