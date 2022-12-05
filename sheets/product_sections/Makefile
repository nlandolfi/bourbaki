# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_sections.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

product_sections.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > product_sections.tex

product_sections.pdf: ../../*.tex ../../trademark.pdf *.tex product_sections.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_sections.tex
	../../latexrun product_sections.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_sections.pdf
	open product_sections.pdf

o:
	make open

terms: product_sections.tex
	bbk_sheet -mode ts product_sections.tex

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
	rm -f product_sections.*
	make reset
	bbk_sheet -mode mk > Makefile