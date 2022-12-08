# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make family_operations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

family_operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > family_operations.tex

family_operations.pdf: ../../*.tex ../../trademark.pdf *.tex family_operations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode family_operations.tex
	../../latexrun family_operations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: family_operations.pdf
	open family_operations.pdf

o:
	make open

terms: family_operations.tex
	bbk_sheet -mode ts family_operations.tex

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
	rm -f family_operations.*
	make reset
	bbk_sheet -mode mk > Makefile