# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make partitions_and_equivalence_relations.pdf

partitions_and_equivalence_relations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partitions_and_equivalence_relations.tex

partitions_and_equivalence_relations.pdf: ../../*.tex ../../trademark.pdf *.tex partitions_and_equivalence_relations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode partitions_and_equivalence_relations.tex
	../../latexrun partitions_and_equivalence_relations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partitions_and_equivalence_relations.pdf
	open partitions_and_equivalence_relations.pdf

o:
	make open

terms: partitions_and_equivalence_relations.tex
	bbk_sheet -mode ts partitions_and_equivalence_relations.tex

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
	rm -f partitions_and_equivalence_relations.*
	make reset
	bbk_sheet -mode mk > Makefile