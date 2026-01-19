document.addEventListener("DOMContentLoaded", function () {
  const pdfForm = document.getElementById("pdf-form");
  const statusText = document.getElementById("status");
  const downloadLink = document.getElementById("download-link");

  pdfForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    e.stopImmediatePropagation();
    e.stopPropagation();

    const formData = new FormData(pdfForm);
    statusText.innerText = "Processing PDF...";
    statusText.className = "text-blue-500 font-bold text-3xl";
    statusText.hidden = false;

    try {
      const response = await fetch("/api/fix-pdf", {
        method: "POST",
        redirect: "error",
        body: formData,
      });

      if (response.ok) {
        const blob = await response.blob();

        const fileUrl = URL.createObjectURL(blob);

        downloadLink.href = fileUrl;
        downloadLink.click();

        statusText.innerText = "PDF Fixed Successfully!";
        statusText.className = "text-green-500 font-bold text-3xl";
        statusText.hidden = false;

        setTimeout(() => URL.revokeObjectURL(fileUrl), 60000);
      } else {
        statusText.innerText = "Error: " + (await response.text());
        statusText.className = "text-red-500 font-bold text-3xl";
        statusText.hidden = false;
      }
    } catch (err) {
      console.log(err);
      statusText.innerText = "Network error occurred.";
      statusText.className = "text-red-500 font-bold text-3xl";
      statusText.hidden = false;
    }
  });
});
