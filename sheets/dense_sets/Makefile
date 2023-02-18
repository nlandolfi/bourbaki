# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dense_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dense_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dense_sets.tex

dense_sets.pdf: ../../*.tex ../../trademark.pdf *.tex dense_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dense_sets.tex
	../../latexrun dense_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dense_sets.pdf
	open dense_sets.pdf

o:
	make open

terms: dense_sets.tex
	bbk_sheet -mode ts dense_sets.tex

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
	rm -f dense_sets.*
	make reset
	bbk_sheet -mode mk > Makefile