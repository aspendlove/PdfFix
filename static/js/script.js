document.addEventListener("DOMContentLoaded", function () {
  const pdfForm = document.getElementById("pdf-form");
  const statusText = document.getElementById("status");

  pdfForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const formData = new FormData(pdfForm);
    statusText.innerText = "Processing PDF...";
    statusText.className = "text-blue-500 font-bold text-3xl";
    statusText.hidden = false;

    try {
      const response = await fetch("/submit", {
        method: "POST",
        body: formData,
      });

      if (response.ok) {
        statusText.innerText = "PDF Fixed Successfully!";
        statusText.className = "text-green-500 font-bold text-3xl";
        statusText.hidden = false;
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
