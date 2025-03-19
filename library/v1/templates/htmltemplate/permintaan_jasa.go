package htmltemplate

type PermintaanJasaData struct {
	Name              string
	PenggunaJasaName  string
	ContractNumber    string
	TanggalPermintaan string
	PRKNumber         string
	NilaiAnggaran     string
	ProjectLocation   string
	PenyediaJasaName  string
	PenugasanName     string
	StartedPenugasan  string
	FinishedPenugasan string
	JenisProyek       string
	ProjectTypeName   string
	Capacity          int64
	CapacityType      string
	Duration          string
	MakerCommentar    string
	ApproverPosition  string
	ApproverName      string
}

//
//	TanggalPenawaran  string
//	TanggalPermintaan string
//	SuratType         string
//	PenggunaJasaName  string
//	PenyediaJasaName  string
//	PenugasanName     string
//	Duration          string
//	TotalMMPlan       string
//	TotalRBTPlan      string
//	MakerCommentar    string
//	ProjectLocations  []string
//	ApproverPosition  string
//	ApproverName      string
//}

const PermintaanJasa = `
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

  .ttd-container>div {
    text-align: center;
  }

  .ttd-container>* {
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

  .content>* {
    margin-bottom: 25px;
  }
</style>

<body>
  <div id="coordinates"></div>
  <div class="sheet-outer A4">
    <section class="sheet padding-5mm">
      <article class="content">
        <div class="text-center">
          <h1> {{.Name}}</h1>
          <h1>{{.PenggunaJasaName}}</h1>
        </div>
        <table>
          <thead>
            <tr>
              <th style="width: 40px">No</th>
              <th style="width: 150px">Item</th>
              <th>Dokumen</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>1</td>
              <td>Dasar Permintaan</td>
              <td>
                Nomor Kontrak {{.ContractNumber}} tanggal
                {{.TanggalPermintaan}}, Nomor PRK {{.PRKNumber}} dengan Nilai Anggaran
                {{.NilaiAnggaran}}
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
              <td>
                {{.ProjectLocation}}
              </td>
            </tr>   
            <tr>
              <td>4</td>
              <td>Para Pihak</td>
              <td>
                <ol>
                  <li>{{.PenyediaJasaName}}</li>
                  <li>{{.PenggunaJasaName}}</li>
                </ol>
              </td>
            </tr>
            <tr>
              <td>5</td>
              <td>Tanggal Permintaan</td>
              <td>{{.TanggalPermintaan}}</td>
            </tr>
            <tr>
              <td>6</td>
              <td>Summary Permintaan</td>
              <td>
                <p>a. Nama Penugasan: {{.PenugasanName}}</p>
                <p>b. Tanggal Mulai Proyek : {{.StartedPenugasan}}</p>
                <p>c. Tanggal Selesai Penugasan : {{.FinishedPenugasan}}</p>
                <p>d. Jenis Proyek :  {{.JenisProyek}}</p>
                <p>e. Tipe Project : {{.ProjectTypeName}}</p>
                <p>f. Kapasitas : {{.Capacity}}</p>
                <p>g. Tipe Kapasitas : {{.CapacityType}}</p>
                <p>h. Jangka Waktu Proyek: {{.Duration}} Bulan</p>
              </td>
            </tr>
          </tbody>
        </table>
        <div>
          <p>Catatan:</p>
          <p>{{.MakerCommentar}}</p>
        </div>

        <div class="ttd-container">
          <div>
            <p style="text-transform: uppercase">{{.ApproverPosition}}</p>
            <div id="signature">
              <div id="speciment"></div>
              <input id="llx" type="hidden" name="llx">
              <input id="lly" type="hidden" name="lly">
              <input id="urx" type="hidden" name="urx">
              <input id="ury" type="hidden" name="ury">
              <input id="page" type="hidden" name="page">
            </div>
            <p style="text-transform: uppercase">{{.ApproverName}}</p>
          </div>
        </div>
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
      const A4_WIDTH_PT = 595.28;  // Points (1/72 of an inch)
      const A4_HEIGHT_PT = 841.89; // Points

      // Calculate scaling factors
      const scaleX = A4_WIDTH_PT / containerRect.width;
      const scaleY = A4_HEIGHT_PT / containerRect.height;

      // Calculate offsets
      const offsetX = containerRect.left;
      const offsetY = containerRect.top;

      // Calculate coordinates relative to the page
      const llx = Math.round((signatureRect.left - offsetX) * scaleX);
      const lly = Math.round(A4_HEIGHT_PT - ((signatureRect.bottom - offsetY) * scaleY));
      const urx = Math.round(llx + (signatureRect.width * scaleX));
      const ury = Math.round(lly + (signatureRect.height * scaleY));

      document.getElementById("llx").value = llx
      document.getElementById("lly").value = lly
      document.getElementById("urx").value = urx
      document.getElementById("ury").value = ury
      document.getElementById("page").value = 1

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
          top: signatureRect.top
        }
      };
    }
    function displayCoordinates() {
      const coords = getSignatureBoxCoordinates();
      const resultDiv = document.getElementById("coordinates");
      resultDiv.innerHTML =
        "<pre>" + JSON.stringify(coords, null, 2) + "</pre>";
    }
    window.addEventListener('load', displayCoordinates);
    window.addEventListener('DOMContentLoaded', displayCoordinates);
  </script>
</body>

</html>

`
