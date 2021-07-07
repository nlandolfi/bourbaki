# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make families.pdf

families.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > families.tex

families.pdf: ../../*.tex ../../trademark.pdf *.tex families.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode families.tex
	../../latexrun families.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: families.pdf
	open families.pdf

o:
	make open

terms: families.tex
	bbk_sheet -mode ts families.tex

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
	rm -f families.*
	make reset
	bbk_sheet -mode mk > Makefile