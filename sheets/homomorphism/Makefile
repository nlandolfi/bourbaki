# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make homomorphism.pdf

homomorphism.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > homomorphism.tex

homomorphism.pdf: ../../*.tex ../../trademark.pdf *.tex homomorphism.tex
	# pdflatex --file-line-error -interaction=nonstopmode homomorphism.tex
	../../latexrun homomorphism.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: homomorphism.pdf
	open homomorphism.pdf

o:
	make open

terms: homomorphism.tex
	bbk_sheet -mode ts homomorphism.tex

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
	rm -f homomorphism.*
	make reset
	bbk_sheet -mode mk > Makefile