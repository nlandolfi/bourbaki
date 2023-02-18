# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make homomorphisms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

homomorphisms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > homomorphisms.tex

homomorphisms.pdf: ../../*.tex ../../trademark.pdf *.tex homomorphisms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode homomorphisms.tex
	../../latexrun homomorphisms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: homomorphisms.pdf
	open homomorphisms.pdf

o:
	make open

terms: homomorphisms.tex
	bbk_sheet -mode ts homomorphisms.tex

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
	rm -f homomorphisms.*
	make reset
	bbk_sheet -mode mk > Makefile