import { sql_query } from "../../lib/db";
import DataTable from "react-data-table-component";

export async function getServerSideProps(ctx) {
  const account = String(ctx.query.address);
  try {
    const data = await sql_query(
      `select * from peer where address="${account}"`
    );
    const datatext = JSON.stringify(data);
    return {
      props: {
        datatext,
      },
    };
  } catch (e) {
    throw Error(e.message);
  }
}

const ToBeClaimedTable = (props) => {
  const columns = [
    {
      name: "HashID",
      selector: (row) => row.hash_id,
    },
    {
      name: "Network",
      selector: (row) => row.network,
    },
    {
      name: "OnlineDuration(s)",
      selector: (row) => row.online_duration,
    },
    {
      name: "CreateTime",
      selector: (row) => row.created_at,
    },
    {
      name: "IsClaimed",
      selector: (row) => row.claimed,
    },
  ];

  return (
    <div>
      <DataTable
        columns={columns}
        data={JSON.parse(props.datatext)}
        pagination
        fixedHeader
      />
    </div>
  );
};

export default ToBeClaimedTable;
