import axios from "axios";

const form: HTMLFormElement = document.getElementById(
  "upload"
) as HTMLFormElement;
form.addEventListener("submit", (e) => {
  e.preventDefault();

  const fileElement: HTMLInputElement = form.elements["upload"];
  if (fileElement.files.length === 0) {
    return;
  }

  const formData = new FormData();
  formData.append("file", fileElement.files[0]);

  axios.post("/api/upload", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
    onUploadProgress: e => {
        console.log(e)
    }
  }).then(_res => {
    const banner = document.getElementById("banner")
    banner.style.display = 'block';
    banner.style.backgroundColor = 'green';
    banner.textContent = "Upload complete."
  })
});
