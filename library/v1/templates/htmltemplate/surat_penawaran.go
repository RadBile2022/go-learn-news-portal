package htmltemplate

type SuratPenawaranData struct {
	TanggalPenawaran  string
	TanggalPermintaan string
	PenggunaJasaName  string
	PenyediaJasaName  string
	PenugasanName     string
	Duration          string
	TotalMMPlan       string
	TotalRBTPlan      string
	MakerCommentar    string
	ProjectLocations  string
	ApproverPosition  string
	ApproverName      string
	BarcodeB64        string
	IsFromAMS         bool
}

const SuratPenawaran = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
  </head>
  <style>
    @import url("https://fonts.googleapis.com/css2?family=Manrope:wght@200..800&display=swap");

    @page {
      size: A4;
      margin: 0;
    }

    *,
    *::after,
    *::before {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: "Manrope", sans-serif;
      font-optical-sizing: auto;
      font-size: 12px;
      margin: 0;
      padding: 0;
    }

    .sheet-outer {
      margin: 0;
    }

    .sheet {
      margin: 0;
      overflow: hidden;
      position: relative;
      box-sizing: border-box;
      page-break-after: always;
    }

    .sheet-outer.A4 .sheet {
      width: 210mm;
      height: 296mm;
    }

    .sheet.padding-5mm {
      padding: 15mm 15mm 15mm 25mm;
    }

    @media screen {
      body {
        background: #e0e0e0;
      }

      .sheet {
        background: white;
        box-shadow: 0 0.5mm 2mm rgba(0, 0, 0, 0.3);
        margin: 5mm auto;
      }
    }

    table {
      width: 100%;
      border-collapse: collapse;
      border-spacing: 0;
      border: 1px solid #d1d5db;
    }

    h1 {
      font-size: 14px;
      font-weight: bold;
      color: #035b71;
    }

    h2 {
      font-size: 12px;
      font-weight: bold;
      color: #035b71;
    }

    .text-left {
      text-align: left;
    }

    .text-center {
      text-align: center;
    }

    table tr {
      border-bottom: 1px solid #d1d5db;
    }

    table thead tr,
    table tfoot tr {
      background-color: #f3f4f6;
    }

    table th {
      text-align: left;
      text-transform: uppercase;
      font-weight: bold;
    }

    table th,
    table td {
      padding: 6px 12px;
    }

    ol {
      padding-left: 16px;
    }

    .ttd-container {
      display: -webkit-box;
      /* wkhtmltopdf uses this one */
      display: -webkit-flex;
      display: flex;
      -webkit-box-pack: end;
      /* wkhtmltopdf uses this one */
      -webkit-justify-content: end;
      justify-content: end;
    }

    .ttd-container > div {
      text-align: center;
    }

    .ttd-container > * {
      margin-bottom: 10px;
    }

    #coordinates {
      display: none;
      position: fixed;
      z-index: 99;
      top: 0;
      left: 0;
      background-color: #f0f0f0;
      opacity: 0.5;
      padding: 10px;
      margin: 10px 0;
      border-radius: 5px;
    }

    @media print {
      .sheet-outer.A4,
      .sheet-outer.A5.landscape {
        width: 210mm;
      }
    }

    #signature {
      margin-top: 10px;
      position: relative;
      margin-bottom: 10px;
    }

    #speciment {
      position: relative;
      width: 100%;
      height: 70px;
    }

    .uppercase {
      text-transform: uppercase;
    }

    .space-y > * {
      margin-bottom: 10px;
    }

    .m-b {
      margin-bottom: 30px;
    }

    .flex {
      display: flex;
    }

    .justify-between {
      justify-content: space-between;
    }

    .items-center {
      align-items: center;
    }

    .content {
      height: 100%;
      width: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }

    .footer {
      display: flex;
      border-top: 2px solid #d1d5db;
      gap: 15px;
      padding: 10px 2px;
      text-align: justify;
    }

    footer > p {
      flex: 1;
    }

    .footer > .barcode {
      height: 40px;
      width: 40px;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  </style>

  <body>
    <div id="coordinates"></div>
    <div class="sheet-outer A4">
      <section class="sheet padding-5mm">
        <article class="content">
          <div class="space-y">
            <div class="text-center space-y m-b">
              <h1 class="uppercase">Surat Penawaran</h1>
              <h1 class="uppercase">{{.PenggunaJasaName}}</h1>
              <h1 class="uppercase">{{.PenugasanName}}</h1>
            </div>
            <p>
              Berikut ini informasi umum Surat Penawaran dari
              <span class="uppercase">{{.PenyediaJasaName}}</span> untuk
              {{.PenugasanName}}:
            </p>
            <table class="m-b">
              <thead>
                <tr>
                  <th style="width: 40px">No</th>
                  <th style="width: 150px">Item</th>
                  <th>Deskripsi</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1</td>
                  <td>Dasar Penawaran</td>
                  <td class="uppercase">
                    {{if .IsFromAMS }} Surat Permintaan Jasa
                    {{.PenggunaJasaName}} Melalui AMS. {{else}} Surat Permintaan
                    Jasa {{.PenggunaJasaName}} Tanggal {{.TanggalPermintaan}}.
                    {{end}}
                  </td>
                </tr>
                <tr>
                  <td>2</td>
                  <td>Nama Penugasan</td>
                  <td>{{.PenugasanName}}</td>
                </tr>
                <tr>
                  <td>3</td>
                  <td>Lokasi</td>
                  <td>{{.ProjectLocations}}</td>
                </tr>
                <tr>
                  <td>4</td>
                  <td>Para Pihak</td>
                  <td>
                    <ol>
                      <li>{{.PenggunaJasaName}}</li>
                      <li>{{.PenyediaJasaName}}</li>
                    </ol>
                  </td>
                </tr>
                <tr>
                  <td>5</td>
                  <td>Tanggal Surat Penawaran</td>
                  <td>{{.TanggalPenawaran}}</td>
                </tr>
                <tr>
                  <td>6</td>
                  <td>Durasi Penugasan</td>
                  <td>{{.Duration}}</td>
                </tr>
                <tr>
                  <td>7</td>
                  <td>Ringkasan Rencana Biaya Tunai dan Rencana MM</td>
                  <td>
                    <p>Rencana MM: {{.TotalMMPlan}} MM</p>
                    <p>Rencana Biaya Tunai: {{.TotalRBTPlan}}</p>
                  </td>
                </tr>
              </tbody>
            </table>
            <p>
              Mohon dapat dievaluasi dan dipastikan ketersediaan anggaran guna
              dilakukan pembahasan lebih lanjut terkait Rencana Penugasan
              tersebut diatas.
            </p>
            <div class="flex justify-between items-center">
              <div>
                <p>Catatan:</p>
                <p>{{.MakerCommentar}}</p>
              </div>

              <div class="ttd-container">
                <div>
                  <p style="text-transform: uppercase">{{.ApproverPosition}}</p>
                  <div id="signature">
                    <div id="speciment"></div>
                    <input id="llx" type="hidden" name="llx" />
                    <input id="lly" type="hidden" name="lly" />
                    <input id="urx" type="hidden" name="urx" />
                    <input id="ury" type="hidden" name="ury" />
                    <input id="page" type="hidden" name="page" />
                  </div>
                  <p style="text-transform: uppercase">{{.ApproverName}}</p>
                </div>
              </div>
            </div>
          </div>
          <footer class="footer">
            <img
              class="barcode"
              src="data:image/png;base64,{{.BarcodeB64}}"
              alt="System Barcode"
            />
            <p>
              Document ini telah ditandatangani secara elektronik menggunakan
              sertifikat elektronik yang diterbitkan oleh Peruri dan menggunakan
              aplikasi SIMPP Ultimate.
            </p>
          </footer>
        </article>
      </section>
    </div>
    <script>
      function getSignatureBoxCoordinates() {
        // Get the sheet container and signature element
        const sheetContainer = document.querySelector(".sheet-outer > .sheet");
        const signatureElement = document.getElementById("signature");
        const signatureBox = document.getElementById("speciment");

        // If elements are not found, return null or default coordinates
        if (!sheetContainer || !signatureElement || !signatureBox) {
          console.error("Required elements not found");
          return null;
        }

        // Get bounding rectangles
        const containerRect = sheetContainer.getBoundingClientRect();
        const signatureRect = signatureBox.getBoundingClientRect();

        // PDF conversion typically uses A4 standard dimensions
        const A4_WIDTH_PT = 595.28; // Points (1/72 of an inch)
        const A4_HEIGHT_PT = 841.89; // Points

        // Calculate scaling factors
        const scaleX = A4_WIDTH_PT / containerRect.width;
        const scaleY = A4_HEIGHT_PT / containerRect.height;

        // Calculate offsets
        const offsetX = containerRect.left;
        const offsetY = containerRect.top;

        // Calculate coordinates relative to the page
        const llx = Math.round((signatureRect.left - offsetX) * scaleX);
        const lly = Math.round(
          A4_HEIGHT_PT - (signatureRect.bottom - offsetY) * scaleY,
        );
        const urx = Math.round(llx + signatureRect.width * scaleX);
        const ury = Math.round(lly + signatureRect.height * scaleY);

        document.getElementById("llx").value = llx;
        document.getElementById("lly").value = lly;
        document.getElementById("urx").value = urx;
        document.getElementById("ury").value = ury;
        document.getElementById("page").value = 1;

        return {
          containerWidth: containerRect.width,
          containerHeight: containerRect.height,
          scaleX,
          scaleY,
          offsetX,
          offsetY,
          llx,
          lly,
          urx,
          ury,
          signatureRect: {
            width: signatureRect.width,
            height: signatureRect.height,
            left: signatureRect.left,
            top: signatureRect.top,
          },
        };
      }
      function displayCoordinates() {
        const coords = getSignatureBoxCoordinates();
        const resultDiv = document.getElementById("coordinates");
        resultDiv.innerHTML =
          "<pre>" + JSON.stringify(coords, null, 2) + "</pre>";
      }
      window.addEventListener("load", displayCoordinates);
      window.addEventListener("DOMContentLoaded", displayCoordinates);
    </script>
  </body>
</html>
`
