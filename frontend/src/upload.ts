import axios from "axios";

const form: HTMLFormElement = document.getElementById(
  "upload"
) as HTMLFormElement;
const progressbar = document.getElementById("progressbar");
const progress = document.getElementById("progress");
const submitBtn = document.getElementById("submit");
const result = document.getElementById("result");
const resultUrl = document.getElementById("resultUrl") as HTMLAnchorElement;

form.addEventListener("submit", (e) => {
  e.preventDefault();

  const fileElement: HTMLInputElement = form.elements["upload"];
  if (fileElement.files.length === 0) {
    return;
  }

  const formData = new FormData();
  formData.append("file", fileElement.files[0]);

  axios
    .post("/api/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
      onUploadProgress: (e) => {
        progressbar.style.display = "block";
        const percent = Math.round((e.loaded * 100) / e.total);
        progress.style.width = `${percent}%`;
        submitBtn.style.display = "none";
      },
    })
    .then((res) => {
      const banner = document.getElementById("banner");
      banner.style.display = "block";
      banner.style.backgroundColor = "green";
      banner.textContent = "Upload complete.";

      submitBtn.style.display = "block";
      progressbar.style.display = "none";
      result.style.display = "block";
      const fileUrl = `${window.location.origin}/file/${res.data.fileId}`;
      resultUrl.href = fileUrl;
      resultUrl.text = fileUrl;
    });
});
