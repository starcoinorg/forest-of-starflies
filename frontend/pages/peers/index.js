import { sql_query } from "../../lib/db";
import DataTable from "react-data-table-component";

export async function getServerSideProps() {
  try {
    const data = await sql_query("select * from peer");
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

const PeerTable = (props) => {
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

export default PeerTable;
