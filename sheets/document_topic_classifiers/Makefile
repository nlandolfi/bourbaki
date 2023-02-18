# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make document_topic_classifiers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

document_topic_classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > document_topic_classifiers.tex

document_topic_classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex document_topic_classifiers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode document_topic_classifiers.tex
	../../latexrun document_topic_classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: document_topic_classifiers.pdf
	open document_topic_classifiers.pdf

o:
	make open

terms: document_topic_classifiers.tex
	bbk_sheet -mode ts document_topic_classifiers.tex

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
	rm -f document_topic_classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile