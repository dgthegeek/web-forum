
const radioButtons = document.querySelectorAll('.topic-checkboxes');
uploadedImageInput =document.querySelector(".post-img")
const checkboxes = document.querySelectorAll('.checkboxes');
const maxLimit = 3; // Set the maximum number of checkboxes allowed to be selected
// Where to display the uploaded image
uploadedImagePlaceholder = document.querySelector(".uploaded-image-placeholder")
uploadedImageInput.addEventListener("input", (e)=>{
  uploadedImagePlaceholder.classList.remove("hidden")
  document.querySelector(".image-btn-uploader").classList.add("hidden")
  document.querySelector(".uploaded-img").src = URL.createObjectURL(e.target.files[0])
})
document.querySelector(".delete-uploaded-image").addEventListener("click", ()=>{

document.querySelector(".image-btn-uploader").classList.remove("hidden")
uploadedImagePlaceholder.classList.add("hidden")
uploadedImageInput.value = ""
})
document.querySelector(".change-uploaded-image").addEventListener("click",()=>{
uploadedImageInput.click()
})


radioButtons.forEach(checkbox => {
  checkbox.addEventListener('change', () => {
    console.log("test")
      let checkedCount = 0;
      radioButtons.forEach(cb => {
          if (cb.checked) {
              checkedCount++;
          document.querySelector(".message").textContent = ""
          }
      });
      if (checkedCount > maxLimit) {
          document.querySelector(".message").textContent = `You cannot select more than   ${maxLimit}  topics`
          checkbox.checked = false; // Prevent checking the checkbox
      }
  });
});