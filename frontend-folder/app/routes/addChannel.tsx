import { Form, redirect, type ActionFunctionArgs } from 'react-router'

export function meta() {
  return [{title: "Add Channel | YT RSS"}];
}

export async function action({request}: ActionFunctionArgs) {
  const formData = await request.formData()
  const handle = formData.get("handle") as string;

  if (!handle) {
    return { error: "Handle required" };
  }

  const res = await fetch(`http://localhost:8080/channel?handleId=${encodeURIComponent(handle as string)}`, {
    method: "PUT",
  })

  if(!res.ok) {
    return { error: "Channel doesn't exist/already added" }
  }

  return redirect("/");
}

export default function addChannel() {
  return (
    <div className="max-w-md mx-auto"> 
      <h1 className="text-2xl font-bold mb-4"> Add Channel</h1>
      <Form method="put" className="space-y-4 bg-white text-black p-4 rounded shadow">
        <div>
	  <label className="block text-gray-700"> YouTube Handle </label>
	  <input type="text" name="handle" 
	    className="border border-gray-300 rounded px-3 py-2 w-full"
	    required />
	</div>
	<button type="submit" className="bg-green-600 text-white px-4 py-2 rounded hover:bg-red-600"> Add Channel </button>
      </Form>
    </div>
  );
}
