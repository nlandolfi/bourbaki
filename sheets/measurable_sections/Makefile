# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make measurable_sections.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

measurable_sections.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > measurable_sections.tex

measurable_sections.pdf: ../../*.tex ../../trademark.pdf *.tex measurable_sections.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode measurable_sections.tex
	../../latexrun measurable_sections.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measurable_sections.pdf
	open measurable_sections.pdf

o:
	make open

terms: measurable_sections.tex
	bbk sheet -mode ts measurable_sections.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measurable_sections.*
	make reset
	bbk sheet -mode mk > Makefile