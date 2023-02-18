# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordering_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

ordering_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > ordering_sets.tex

ordering_sets.pdf: ../../*.tex ../../trademark.pdf *.tex ordering_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordering_sets.tex
	../../latexrun ordering_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordering_sets.pdf
	open ordering_sets.pdf

o:
	make open

terms: ordering_sets.tex
	bbk_sheet -mode ts ordering_sets.tex

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
	rm -f ordering_sets.*
	make reset
	bbk_sheet -mode mk > Makefile