# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complex_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_measures.tex

complex_measures.pdf: ../../*.tex ../../trademark.pdf *.tex complex_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_measures.tex
	../../latexrun complex_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_measures.pdf
	open complex_measures.pdf

o:
	make open

terms: complex_measures.tex
	bbk_sheet -mode ts complex_measures.tex

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
	rm -f complex_measures.*
	make reset
	bbk_sheet -mode mk > Makefile