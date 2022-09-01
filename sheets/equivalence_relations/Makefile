# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make equivalence_relations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

equivalence_relations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > equivalence_relations.tex

equivalence_relations.pdf: ../../*.tex ../../trademark.pdf *.tex equivalence_relations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode equivalence_relations.tex
	../../latexrun equivalence_relations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equivalence_relations.pdf
	open equivalence_relations.pdf

o:
	make open

terms: equivalence_relations.tex
	bbk_sheet -mode ts equivalence_relations.tex

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
	rm -f equivalence_relations.*
	make reset
	bbk_sheet -mode mk > Makefile