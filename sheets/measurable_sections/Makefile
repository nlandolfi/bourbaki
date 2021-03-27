# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measurable_sections.pdf
	make graph.pdf

measurable_sections.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measurable_sections.tex

measurable_sections.pdf: ../../*.tex ../../trademark.pdf *.tex measurable_sections.tex
	pdflatex --file-line-error -interaction=nonstopmode measurable_sections.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measurable_sections.pdf
	open measurable_sections.pdf

o:
	make open

defs: measurable_sections.tex
	defs measurable_sections.tex

spell:
	aspell -c measurable_sections.tex

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
	bbk_sheet -mode mk > Makefile