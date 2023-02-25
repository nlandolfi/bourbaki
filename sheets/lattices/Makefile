# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make lattices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

lattices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > lattices.tex

lattices.pdf: ../../*.tex ../../trademark.pdf *.tex lattices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode lattices.tex
	../../latexrun lattices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: lattices.pdf
	open lattices.pdf

o:
	make open

terms: lattices.tex
	bbk_sheet -mode ts lattices.tex

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
	rm -f lattices.*
	make reset
	bbk_sheet -mode mk > Makefile